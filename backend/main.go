package main

import (
	"backend/app/db/postgre"
	"backend/migrations"
	"backend/pkg/api"
	"log"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	db, err := postgre.ConnectDB()
	if err != nil {
		log.Fatal("Could not connect to Postgres:", err)
	}

	err = migrations.RunMigrations(db)
	if err != nil {
		log.Fatal("Could not run migrations:", err)
	}

	api.Start(db)
}
