package migrations

import (
	"gorm.io/gorm"
)

func init() {
	Migrations = append(Migrations, Migration{
		Name: "20250327165812_create_term_deposits_table",
		Up: func(db *gorm.DB) error {
			err := db.Exec(`
				CREATE TABLE term_deposits (
					id SERIAL PRIMARY KEY,
					bank_account_id INT REFERENCES bank_accounts(id),
					amount DECIMAL(15, 2) NOT NULL,
					interest_rate DECIMAL(5, 2) NOT NULL,
					term_months INT NOT NULL,
					start_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
					maturity_date TIMESTAMP,
					status VARCHAR(20) DEFAULT 'active'
				);`).Error
			if err != nil {
				return err
			}

			err = db.Exec(`
				INSERT INTO term_deposits (bank_account_id, amount, interest_rate, term_months, start_date, maturity_date, status) VALUES
				(1, 1000000, 0.05, 6, '2025-03-27', '2025-09-27', 'active'),
				(2, 2000000, 0.06, 12, '2025-03-27', '2026-03-27', 'active'),
				(3, 3000000, 0.07, 24, '2025-03-27', '2027-03-27', 'active'),
				(4, 4000000, 0.08, 36, '2025-03-27', '2028-03-27', 'active'),
				(5, 5000000, 0.09, 48, '2025-03-27', '2029-03-27', 'active'),
				(6, 6000000, 0.10, 60, '2025-03-27', '2030-03-27', 'active'),
				(7, 7000000, 0.11, 72, '2025-03-27', '2031-03-27', 'active'),
				(8, 8000000, 0.12, 84, '2025-03-27', '2032-03-27', 'active'),
				(9, 9000000, 0.13, 96, '2025-03-27', '2033-03-27', 'active'),
				(10, 10000000, 0.14, 108, '2025-03-27', '2034-03-27', 'active')
			`).Error
			if err != nil {
				return err
			}

			return nil
		},
		Down: func(db *gorm.DB) error {
			err := db.Exec("DROP TABLE IF EXISTS term_deposits").Error
			if err != nil {
				return err
			}
			return nil
		},
	})
}
