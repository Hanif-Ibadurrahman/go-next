package migrations

import (
	"backend/app/util"
	"backend/pkg/api/v1/user/models"

	"gorm.io/gorm"
)

func init() {
	Migrations = append(Migrations, Migration{
		Name: "20250327141421_create_users_table",
		Up: func(db *gorm.DB) error {
			err := db.Exec(`
				CREATE TABLE IF NOT EXISTS users (
					id_user SERIAL PRIMARY KEY,
					name VARCHAR(50),
					username VARCHAR(255) UNIQUE,
					password VARCHAR(255),
					level INTEGER,
					created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
					updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
					last_login TIMESTAMP,
					is_active INTEGER,
					deleted_at TIMESTAMP
				)`).Error
			if err != nil {
				return err
			}
			err = db.Exec(`
				CREATE INDEX IF NOT EXISTS idx_users_deleted_at ON users (deleted_at)
			`).Error
			if err != nil {
				return err
			}
			// Seed data
			users := []models.User{
				{
					IDUser:   1,
					Name:     "Customer 1",
					Username: "customer1",
					Password: util.HashPassword("password123"),
					Level:    2,
					IsActive: 1,
				},
				{
					IDUser:   2,
					Name:     "Customer 2",
					Username: "customer2",
					Password: util.HashPassword("password123"),
					Level:    2,
					IsActive: 1,
				},
				{
					IDUser:   3,
					Name:     "Customer 3",
					Username: "customer3",
					Password: util.HashPassword("password123"),
					Level:    2,
					IsActive: 1,
				},
				{
					IDUser:   4,
					Name:     "Customer 4",
					Username: "customer4",
					Password: util.HashPassword("password123"),
					Level:    2,
					IsActive: 1,
				},
				{
					IDUser:   5,
					Name:     "Customer 5",
					Username: "customer5",
					Password: util.HashPassword("password123"),
					Level:    2,
					IsActive: 1,
				},
			}

			for _, user := range users {
				if err := db.Create(&user).Error; err != nil {
					return err
				}
			}
			return nil
		},
		Down: func(db *gorm.DB) error {
			err := db.Exec("DROP TABLE IF EXISTS users").Error
			if err != nil {
				return err
			}
			return nil
		},
	})
}
