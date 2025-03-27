package migrations

import (
	"gorm.io/gorm"
)

func init() {
	Migrations = append(Migrations, Migration{
		Name: "20250327165346_create_bank_accounts_table",
		Up: func(db *gorm.DB) error {
			err := db.Exec(`
				CREATE TABLE bank_accounts (
					id SERIAL PRIMARY KEY,
					user_id INT REFERENCES users(id_user),
					account_number VARCHAR(20) UNIQUE NOT NULL,
					balance DECIMAL(15, 2) DEFAULT 0.00,
					created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
				);`).Error
			if err != nil {
				return err
			}
			err = db.Exec(`
				INSERT INTO bank_accounts (user_id, account_number, balance) VALUES
				(1, '1234567890', 1000000),
				(1, '2345678901', 2000000),
				(2, '3456789012', 3000000),
				(2, '4567890123', 4000000),
				(3, '5678901234', 5000000),
				(3, '6789012345', 6000000),
				(4, '7890123456', 7000000),
				(4, '8901234567', 8000000),
				(5, '9012345678', 9000000),
				(5, '0123456789', 10000000)
			`).Error
			if err != nil {
				return err
			}
			return nil
		},
		Down: func(db *gorm.DB) error {
			err := db.Exec("DROP TABLE IF EXISTS bank_accounts").Error
			if err != nil {
				return err
			}
			return nil
		},
	})
}
