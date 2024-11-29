package main

import (
	"BuildasTechnicalAssessmentGo/internal/config"
	"BuildasTechnicalAssessmentGo/internal/database"
	"BuildasTechnicalAssessmentGo/internal/repository"
	"BuildasTechnicalAssessmentGo/pkg/handlers"
	"BuildasTechnicalAssessmentGo/pkg/services"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
)

// The main is the entry point of the program.
//
// It performs the following steps:
// 1. Loads environment variables.
// 2. Initializes the database and runs migrations.
// 3. Create a new Gin router.
// 4. Loads HTML templates.
// 5. Serves static files for assets (CSS, JS, etc.).
// 6. Register routes.
// 7. Start the server.
func main() {
	// Load environment variables
	cfg := config.LoadConfig()

	// Initialize the database and run migrations
	database.ConnectDB(cfg)

	// Create a new Gin router
	r := gin.New()

	// Set up the repository and service layer
	userRepo := repository.PostgresUserRepository{DB: database.DB}
	userService := services.UserService{Repo: &userRepo}

	// Get the absolute base path of the project
	basePath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// Construct paths for templates and static assets
	templatesPath := filepath.Join(basePath, "web", "templates", "*")
	assetsPath := filepath.Join(basePath, "web", "assets")

	// Load HTML templates
	r.LoadHTMLGlob(templatesPath)

	// Serve static files for assets (CSS, JS, etc.)
	r.Static("/assets", assetsPath)

	// Register routes
	handlers.RegisterRoutes(r, &userService)

	// Start the server
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
