package postgre

import (
	"backend/app/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	dbUser := config.GetConfig().PostgreDBUser
	dbPassword := config.GetConfig().PostgreDBPassword
	dbHost := config.GetConfig().PostgreDBHost
	dbName := config.GetConfig().PostgreDBName
	dbPort := config.GetConfig().PostgreDBPort

	// Construct DSN (Data Source Name)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort)

	// Open the database connection
	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Database connection established")

	return db, nil
}
