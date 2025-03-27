package migrations

import (
	"gorm.io/gorm"
)

func init() {
	Migrations = append(Migrations, Migration{
		Name: "20250327165632_create_pockets_table",
		Up: func(db *gorm.DB) error {
			err := db.Exec(`
				CREATE TABLE pockets (
					id SERIAL PRIMARY KEY,
					bank_account_id INT REFERENCES bank_accounts(id),
					pocket_name VARCHAR(50) NOT NULL,
					balance DECIMAL(15, 2) DEFAULT 0.00,
					created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
				);`).Error
			if err != nil {
				return err
			}
			err = db.Exec(`
				INSERT INTO pockets (bank_account_id, pocket_name, balance) VALUES
				(1, 'Pocket 1', 1000000),
				(2, 'Pocket 2', 2000000),
				(3, 'Pocket 3', 3000000),
				(4, 'Pocket 4', 4000000),
				(5, 'Pocket 5', 5000000),
				(6, 'Pocket 6', 6000000),
				(7, 'Pocket 7', 7000000),
				(8, 'Pocket 8', 8000000),
				(9, 'Pocket 9', 9000000),
				(10, 'Pocket 10', 10000000)
			`).Error
			if err != nil {
				return err
			}
			return nil
		},
		Down: func(db *gorm.DB) error {
			err := db.Exec("DROP TABLE IF EXISTS pockets").Error
			if err != nil {
				return err
			}
			return nil
		},
	})
}
