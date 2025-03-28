package impl

import (
	"backend/pkg/api/v1/user/models"
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock db: %v", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		t.Fatalf("failed to open gorm db: %v", err)
	}

	return gormDB, mock
}

func TestSearch(t *testing.T) {
	gormDB, mock := setupMockDB(t)
	userRepo := NewUser(gormDB)
	ctx := context.Background()

	tests := []struct {
		name      string
		req       models.QuerySearch
		setupMock func(mock sqlmock.Sqlmock)
		wantLen   int
		wantErr   bool
		checkFunc func(t *testing.T, users []models.UserDetail)
	}{
		{
			name: "Search by username",
			req:  models.QuerySearch{Q: "user1"},
			setupMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{
					"id_user", "name", "username", "level", "created_at", "updated_at", "last_login", "is_active", "account_number", "balance",
				}).AddRow(
					1, "User One", "user1", 1, time.Now(), time.Now(), nil, 1, "ACC001", 1000.00,
				)
				mock.ExpectQuery(`SELECT users\.id_user, users\.name, users\.username, users\.level, users\.created_at, users\.updated_at, users\.last_login, users\.is_active, bank_accounts\.account_number, bank_accounts\.balance FROM "users" LEFT JOIN bank_accounts ON bank_accounts\.user_id = users\.id_user LEFT JOIN pockets ON pockets\.bank_account_id = bank_accounts\.id LEFT JOIN term_deposits ON term_deposits\.bank_account_id = bank_accounts\.id WHERE users\.name ILIKE \$1 OR users\.username ILIKE \$2 OR bank_accounts\.account_number ILIKE \$3 OR pockets\.pocket_name ILIKE \$4 OR term_deposits\.status ILIKE \$5`).
					WithArgs("%user1%", "%user1%", "%user1%", "%user1%", "%user1%").
					WillReturnRows(rows)

				pocketRows := sqlmock.NewRows([]string{"id", "bank_account_id", "pocket_name", "balance", "created_at"}).
					AddRow(1, 1, "Savings", 500.00, time.Now())
				mock.ExpectQuery(`SELECT \* FROM "pockets" WHERE bank_account_id IN \(SELECT id FROM bank_accounts WHERE user_id = \$1\) AND pocket_name ILIKE \$2`).
					WithArgs(1, "%user1%").
					WillReturnRows(pocketRows)

				termRows := sqlmock.NewRows([]string{"id", "bank_account_id", "amount", "interest_rate", "term_months", "start_date", "maturity_date", "status"}).
					AddRow(1, 1, 2000.00, 3.5, 12, time.Now(), time.Now().AddDate(1, 0, 0), "active")
				mock.ExpectQuery(`SELECT \* FROM "term_deposits" WHERE bank_account_id IN \(SELECT id FROM bank_accounts WHERE user_id = \$1\) AND status ILIKE \$2`).
					WithArgs(1, "%user1%").
					WillReturnRows(termRows)
			},
			wantLen: 1,
			wantErr: false,
			checkFunc: func(t *testing.T, users []models.UserDetail) {
				if len(users) == 0 {
					t.Errorf("expected users, got none")
					return
				}
				assert.Equal(t, "user1", users[0].Username)
				assert.Equal(t, "ACC001", users[0].AccountNumber)
				assert.Len(t, users[0].Pockets, 1)
				assert.Equal(t, "Savings", users[0].Pockets[0].PocketName)
				assert.Len(t, users[0].TermDeposits, 1)
			},
		},
		{
			name: "Search with empty query",
			req:  models.QuerySearch{Q: ""},
			setupMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{
					"id_user", "name", "username", "level", "created_at", "updated_at", "last_login", "is_active", "account_number", "balance",
				}).
					AddRow(1, "User One", "user1", 1, time.Now(), time.Now(), nil, 1, "ACC001", 1000.00).
					AddRow(2, "User Two", "user2", 2, time.Now(), time.Now(), nil, 1, "ACC002", 2000.00)
				mock.ExpectQuery(`SELECT users\.id_user, users\.name, users\.username, users\.level, users\.created_at, users\.updated_at, users\.last_login, users\.is_active, bank_accounts\.account_number, bank_accounts\.balance FROM "users" LEFT JOIN bank_accounts ON bank_accounts\.user_id = users\.id_user LEFT JOIN pockets ON pockets\.bank_account_id = bank_accounts\.id LEFT JOIN term_deposits ON term_deposits\.bank_account_id = bank_accounts\.id`).
					WillReturnRows(rows)

				pocketRows1 := sqlmock.NewRows([]string{"id", "bank_account_id", "pocket_name", "balance", "created_at"}).
					AddRow(1, 1, "Savings", 500.00, time.Now())
				mock.ExpectQuery(`SELECT \* FROM "pockets" WHERE bank_account_id IN \(SELECT id FROM bank_accounts WHERE user_id = \$1\)`).
					WithArgs(1).
					WillReturnRows(pocketRows1)

				termRows1 := sqlmock.NewRows([]string{"id", "bank_account_id", "amount", "interest_rate", "term_months", "start_date", "maturity_date", "status"}).
					AddRow(1, 1, 2000.00, 3.5, 12, time.Now(), time.Now().AddDate(1, 0, 0), "active")
				mock.ExpectQuery(`SELECT \* FROM "term_deposits" WHERE bank_account_id IN \(SELECT id FROM bank_accounts WHERE user_id = \$1\)`).
					WithArgs(1).
					WillReturnRows(termRows1)

				pocketRows2 := sqlmock.NewRows([]string{"id", "bank_account_id", "pocket_name", "balance", "created_at"}).
					AddRow(2, 2, "Emergency", 1000.00, time.Now())
				mock.ExpectQuery(`SELECT \* FROM "pockets" WHERE bank_account_id IN \(SELECT id FROM bank_accounts WHERE user_id = \$1\)`).
					WithArgs(2).
					WillReturnRows(pocketRows2)

				termRows2 := sqlmock.NewRows([]string{"id", "bank_account_id", "amount", "interest_rate", "term_months", "start_date", "maturity_date", "status"}).
					AddRow(2, 2, 3000.00, 4.0, 24, time.Now(), time.Now().AddDate(2, 0, 0), "active")
				mock.ExpectQuery(`SELECT \* FROM "term_deposits" WHERE bank_account_id IN \(SELECT id FROM bank_accounts WHERE user_id = \$1\)`).
					WithArgs(2).
					WillReturnRows(termRows2)
			},
			wantLen: 2,
			wantErr: false,
			checkFunc: func(t *testing.T, users []models.UserDetail) {
				assert.Len(t, users, 2)
				if len(users) != 2 {
					t.Errorf("expected 2 users, got %d", len(users))
					return
				}
				assert.Equal(t, "user1", users[0].Username)
				assert.Equal(t, "user2", users[1].Username)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock(mock)
			users, err := userRepo.Search(ctx, tt.req)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Len(t, users, tt.wantLen)
			if tt.checkFunc != nil {
				tt.checkFunc(t, users)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
