package impl

import (
	"backend/pkg/api/v1/user/models"
	"context"

	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{db: db}
}

func (u *User) Search(ctx context.Context, req models.QuerySearch) ([]models.UserDetail, error) {
	var flatUsers []models.FlatUser
	query := u.db.WithContext(ctx).
		Table("users").
		Select(`
            users.id_user,
            users.name,
            users.username,
            users.level,
            users.created_at,
            users.updated_at,
            users.last_login,
            users.is_active,
            bank_accounts.account_number,
            bank_accounts.balance
        `).
		Joins("LEFT JOIN bank_accounts ON bank_accounts.user_id = users.id_user").
		Joins("LEFT JOIN pockets ON pockets.bank_account_id = bank_accounts.id").
		Joins("LEFT JOIN term_deposits ON term_deposits.bank_account_id = bank_accounts.id")

	if req.Q != "" {
		query = query.Where(`
            users.name ILIKE ? OR 
            users.username ILIKE ? OR 
            bank_accounts.account_number ILIKE ? OR 
            pockets.pocket_name ILIKE ? OR 
            term_deposits.status ILIKE ?
        `, "%"+req.Q+"%", "%"+req.Q+"%", "%"+req.Q+"%", "%"+req.Q+"%", "%"+req.Q+"%")
	}

	err := query.Scan(&flatUsers).Error
	if err != nil {
		return nil, err
	}

	seenUsers := make(map[int]bool)
	uniqueUsers := []models.UserDetail{}

	for _, fu := range flatUsers {
		if !seenUsers[fu.IDUser] {
			user := models.UserDetail{
				IDUser:        fu.IDUser,
				Name:          fu.Name,
				Username:      fu.Username,
				Level:         fu.Level,
				CreatedAt:     fu.CreatedAt,
				UpdatedAt:     fu.UpdatedAt,
				LastLogin:     fu.LastLogin,
				IsActive:      fu.IsActive,
				AccountNumber: fu.AccountNumber,
				Balance:       fu.Balance,
			}

			// Pockets
			var pockets []models.Pocket
			pocketQuery := u.db.WithContext(ctx).
				Table("pockets").
				Where("bank_account_id IN (SELECT id FROM bank_accounts WHERE user_id = ?)", user.IDUser)
			if req.Q != "" {
				pocketQuery = pocketQuery.Where("pocket_name ILIKE ?", "%"+req.Q+"%")
			}
			err = pocketQuery.Find(&pockets).Error
			if err != nil {
				return nil, err
			}
			user.Pockets = pockets

			// Term Deposits
			var termDeposits []models.TermDeposit
			termQuery := u.db.WithContext(ctx).
				Table("term_deposits").
				Where("bank_account_id IN (SELECT id FROM bank_accounts WHERE user_id = ?)", user.IDUser)
			if req.Q != "" {
				termQuery = termQuery.Where("status ILIKE ?", "%"+req.Q+"%")
			}
			err = termQuery.Find(&termDeposits).Error
			if err != nil {
				return nil, err
			}
			user.TermDeposits = termDeposits

			seenUsers[fu.IDUser] = true
			uniqueUsers = append(uniqueUsers, user)
		}
	}

	return uniqueUsers, nil
}

func (u *User) CreateUser(ctx context.Context, req models.User) (*models.User, error) {
	query := `
		INSERT INTO users (username, name, password, level, is_active) 
		VALUES (?, ?, ?, ?, ?)
	`
	result := u.db.WithContext(ctx).Exec(query, req.Username, req.Name, req.Password, req.Level, req.IsActive)
	if result.Error != nil {
		return nil, result.Error
	}

	// Fetch inserted user
	var user models.User
	err := u.db.WithContext(ctx).
		Raw("SELECT id_user, username, name, level FROM users WHERE username = ?", req.Username).
		Scan(&user).Error

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) UpdateUser(ctx context.Context, req models.User) error {
	query := "UPDATE users SET "
	args := []interface{}{}

	if req.Name != "" {
		query += "name = ?,"
		args = append(args, req.Name)
	}

	if req.Password != "" {
		query += "password = ?,"
		args = append(args, req.Password)
	}

	if len(args) == 0 {
		return nil
	}

	query = query[:len(query)-1] // Remove trailing comma
	query += " WHERE id_user = ?"
	args = append(args, req.IDUser)

	result := u.db.WithContext(ctx).Exec(query, args...)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (u *User) DeleteUser(ctx context.Context, userId int) error {
	query := `
		UPDATE users SET
			is_active = 0
		WHERE
			id_user = ?`
	result := u.db.WithContext(ctx).Exec(query, userId)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *User) UsernameExists(ctx context.Context, username string) (bool, error) {
	var count int64
	err := u.db.WithContext(ctx).
		Raw("SELECT COUNT(*) FROM users WHERE username = ?", username).
		Scan(&count).Error

	if err != nil {
		return false, err
	}
	return count > 0, nil
}
