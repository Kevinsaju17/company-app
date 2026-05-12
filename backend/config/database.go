package config

import (
	"fmt"
	"log"

	"company-app/backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


// Global database variable
// Accessible from anywhere in project
var DB *gorm.DB


// ConnectDatabase establishes connection with PostgreSQL
func ConnectDatabase() {

	// DSN = Data Source Name
	// Contains DB connection details
	//
	// IMPORTANT:
	// Replace password with YOUR postgres password
	dsn := "host=localhost user=postgres password=postgres123 dbname=company_app port=5432 sslmode=disable"

	var err error

	// Open PostgreSQL connection
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// If connection fails
	if err != nil {

		// Stop application immediately
		log.Fatal("Database connection failed")
	}

	fmt.Println("Database connected successfully")


	// AutoMigrate automatically creates/updates tables
	// based on Go structs/models
	DB.AutoMigrate(&models.User{})
}