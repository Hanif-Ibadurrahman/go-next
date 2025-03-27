package migrations

import (
	"sort"

	"gorm.io/gorm"
)

type Migration struct {
	Name string
	Up   func(db *gorm.DB) error
	Down func(db *gorm.DB) error
}

var Migrations []Migration

// RunMigrations executes all registered migrations in order
func RunMigrations(db *gorm.DB) error {

	// Ensure the migrations table exists before proceeding
	if err := ensureMigrationsTable(db); err != nil {
		return err
	}

	// Sort migrations by name (timestamp order)
	sort.Slice(Migrations, func(i, j int) bool {
		return Migrations[i].Name < Migrations[j].Name
	})

	// Run only unapplied migrations
	for _, m := range Migrations {
		if err := runMigrationIfNotApplied(db, m); err != nil {
			return err
		}
	}
	return nil
}

// RollbackMigrations rolls back all registered migrations in reverse order
func RollbackMigrations(db *gorm.DB) error {
	// Sort migrations in reverse order
	sort.Slice(Migrations, func(i, j int) bool {
		return Migrations[i].Name > Migrations[j].Name
	})

	// Roll back only applied migrations
	for _, m := range Migrations {
		var count int64
		err := db.Raw("SELECT COUNT(*) FROM migrations WHERE name = ?", m.Name).Scan(&count).Error
		if err != nil {
			return err
		}
		if count > 0 { // Migration was applied
			err := m.Down(db)
			if err != nil {
				return err
			}
			err = db.Exec("DELETE FROM migrations WHERE name = ?", m.Name).Error
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// ensureMigrationsTable creates the migrations table if it doesn’t exist
func ensureMigrationsTable(db *gorm.DB) error {
	// Check if the migrations table exists by attempting a simple query
	var count int64
	err := db.Raw("SELECT COUNT(*) FROM migrations LIMIT 1").Scan(&count).Error
	if err != nil && err.Error() == `ERROR: relation "migrations" does not exist (SQLSTATE 42P01)` {
		// Table doesn’t exist, so create it
		err = db.Exec(`
			CREATE TABLE IF NOT EXISTS migrations (
				name VARCHAR(255) PRIMARY KEY,
				applied_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
			)
		`).Error
		if err != nil {
			return err
		}
	} else if err != nil {
		// Some other error occurred
		return err
	}
	return nil
}

func runMigrationIfNotApplied(db *gorm.DB, m Migration) error {
	var count int64
	err := db.Raw("SELECT COUNT(*) FROM migrations WHERE name = ?", m.Name).Scan(&count).Error
	if err != nil {
		return err
	}

	if count == 0 { // Migration hasn’t been applied
		err = m.Up(db)
		if err != nil {
			return err
		}
		// Record the migration as applied
		err = db.Exec("INSERT INTO migrations (name) VALUES (?)", m.Name).Error
		if err != nil {
			return err
		}
	}
	return nil
}
