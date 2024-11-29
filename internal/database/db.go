package database

import (
	"BuildasTechnicalAssessmentGo/internal/config"
	"BuildasTechnicalAssessmentGo/pkg/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

// ConnectDB connects to the PostgresSQL database, migrates the User model,
// and logs a success message.
//
// This function takes the application configuration as a parameter and
// uses it to construct the PostgresSQL connection string. It then opens
// a connection to the database using GORM and migrates the User model.
// If either of these operations fails, it logs a fatal error message.
// Otherwise, it logs a success message.
func ConnectDB(cfg config.Config) {
	// Build the connection string
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.PostgresHost, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB, cfg.PostgresPort,
	)

	// Open a connection to the database
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Migrate the User model
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Log a success message
	log.Println("Database connection and migration successful")
}
