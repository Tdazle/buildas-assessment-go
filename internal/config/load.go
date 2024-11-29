package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// LoadConfig loads environment variables from a .env file and returns a
// Config struct with the application configuration.
//
// It first loads the .env file using dotenv, and logs a fatal error
// if there is any issue loading the file. Then, it populates the Config
// struct with the environment variables and returns it.
//
// The environment variables are loaded from the .env file using the
// following names:
// - POSTGRES_USER
// - POSTGRES_PASSWORD
// - POSTGRES_DB
// - POSTGRES_HOST
// - POSTGRES_PORT
func LoadConfig() Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Populate the Config struct with environment variables
	return Config{
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresDB:       os.Getenv("POSTGRES_DB"),
		PostgresHost:     os.Getenv("POSTGRES_HOST"),
		PostgresPort:     os.Getenv("POSTGRES_PORT"),
	}
}
