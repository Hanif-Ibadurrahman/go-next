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
					Name:     "User 1",
					Username: "user1",
					Password: util.HashPassword("password123"),
					Level:    2,
					IsActive: 1,
				},
				{
					IDUser:   2,
					Name:     "User 2",
					Username: "user2",
					Password: util.HashPassword("password123"),
					Level:    2,
					IsActive: 1,
				},
				{
					IDUser:   3,
					Name:     "User 3",
					Username: "user3",
					Password: util.HashPassword("password123"),
					Level:    2,
					IsActive: 1,
				},
				{
					IDUser:   4,
					Name:     "User 4",
					Username: "user4",
					Password: util.HashPassword("password123"),
					Level:    2,
					IsActive: 1,
				},
				{
					IDUser:   5,
					Name:     "User 5",
					Username: "user5",
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
