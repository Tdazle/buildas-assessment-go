package config

// Config holds the application configuration (e.g., database settings)
type Config struct {
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
	PostgresHost     string
	PostgresPort     string
}
