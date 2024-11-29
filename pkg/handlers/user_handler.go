package handlers

import (
	"BuildasTechnicalAssessmentGo/pkg/middlewares"
	"BuildasTechnicalAssessmentGo/pkg/services"
	"BuildasTechnicalAssessmentGo/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
)

// RegisterRoutes registers all the routes for the user handlers.
//
// The routes are registered with the given gin.Engine. The routes are divided
// into two groups: one for the user routes, and one for the routes that require
// authentication. The routes in the authUser group are protected by the
// middleware.AuthMiddleware function, which checks for a valid JWT token in
// the Authorization header of the request. The userService parameter is used
// to create a new instance of the UserService, which is used by the handlers
// to interact with the database.
func RegisterRoutes(r *gin.Engine, userService services.UserServiceInterface) {
	user := r.Group("/api/v1/user")
	authUser := r.Group("/api/v1/user")
	authUser.Use(middlewares.AuthMiddleware())
	{
		user.GET("/register", UserRegisterForm)
		user.POST("/register", func(c *gin.Context) { RegisterUser(c, userService) })
		user.GET("/login", LoginUserForm)
		user.POST("/login", func(c *gin.Context) { LoginUser(c, userService) })
		authUser.GET("/home", func(c *gin.Context) { Home(c, userService) })
		authUser.POST("/add", func(c *gin.Context) { AddUser(c, userService) })
	}
}

// Home handles the HTTP GET request for the home page.
//
// It extracts the username from the JWT claims stored in the request context,
// fetches all users from the database using the provided userService,
// and renders the home.html template with the user's username and a list of users.
//
// If the user list is not fetched successfully, it responds with HTTP status 500
// and an error message. Otherwise, it responds with HTTP status 200 and the
// rendered template.
func Home(context *gin.Context, userService services.UserServiceInterface) {
	// Get the username from the claims
	claims := context.MustGet("claims").(*jwt.MapClaims)
	username := (*claims)["username"]

	// Fetch all users from the database
	users, err := userService.GetAllUsers()
	if err != nil {
		context.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": "Failed to load users"})
		return
	}

	// Render the home page with a user list
	context.HTML(http.StatusOK, "home.html", gin.H{
		"username": username,
		"users":    users,
	})
}

// AddUser handles the HTTP POST request for adding a new user.
//
// It extracts the username and password from the form data,
// calls the userService to register the new user,
// and redirects back to the home page if successful.
//
// If the registration fails, it responds with HTTP status 500
// and an error message.
func AddUser(context *gin.Context, userService services.UserServiceInterface) {
	//Get post-data
	username := context.PostForm("username")
	password := context.PostForm("password")

	// Call the user service to register the new user
	err := userService.RegisterUser(username, password)
	if err != nil {
		context.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": err.Error()})
		return
	}

	// Redirect back to the home page
	context.Redirect(http.StatusSeeOther, "/api/v1/user/home")
}

// LoginUserForm handles the HTTP GET request for the login form.
//
// It checks for an existing Authorization cookie. If it exists, it redirects
// to /home. If not, it renders the login.html template with an empty
// gin.H object.
func LoginUserForm(context *gin.Context) {
	// Use the helper function to check for redirection
	if middlewares.RedirectIfAuthenticated(context) {
		return
	}

	// If no token, render the login page
	context.HTML(http.StatusOK, "login.html", gin.H{})
}

// LoginUser handles the HTTP POST request for user login.
//
// It extracts the username and password from the form data,
// calls the userService to validate the credentials,
// generates a JWT token using the utils helper function,
// sets the token as a cookie (optional but useful for session management),
// and renders the home.html file with the user's username.
//
// If the credentials are invalid, it responds with HTTP status 401
// and an error message. If there is an error generating the token,
// it responds with HTTP status 500 and an error message.
func LoginUser(context *gin.Context, userService services.UserServiceInterface) {
	// Extract credentials from the request
	username := context.PostForm("username")
	password := context.PostForm("password")

	// Validate user credentials
	user, err := userService.GetUserByUsername(username)
	if err != nil || user == nil {
		context.HTML(http.StatusUnauthorized, "error.html", gin.H{"error": "Invalid credentials"})
		return
	}

	// Check password
	if err := userService.CheckPassword(user.Password, password); err != nil {
		context.HTML(http.StatusUnauthorized, "error.html", gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token using the helper function
	token, err := utils.GenerateJWT(user) // Call the helper function from utils
	if err != nil {
		context.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": "Could not generate token"})
		return
	}

	// Set the token as a cookie (optional but useful for session management)
	context.SetCookie("Authorization", token, 3600, "/", "localhost", false, true)

	// Render the home.html file
	context.HTML(http.StatusOK, "home.html", gin.H{
		"username": user.Username,
	})

}

// RegisterUser handles the HTTP POST request for user registration.
//
// It extracts the username and password from the form data,
// calls the userService to manage the registration logic,
// and returns a JSON response indicating the registration outcome.
//
// If registration is successful, it responds with HTTP status 200
// and a success message. If an error occurs, it responds with
// HTTP status 400 and an error message.
func RegisterUser(context *gin.Context, userService services.UserServiceInterface) {
	// Extract credentials from the request
	username := context.PostForm("username")
	password := context.PostForm("password")

	// Call the userService to handle registration logic
	err := userService.RegisterUser(username, password)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate JWT token after successful registration
	user, _ := userService.GetUserByUsername(username)
	token, err := utils.GenerateJWT(user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Set the token as a cookie
	context.SetCookie("Authorization", token, 3600, "/", "localhost", false, true)

	// Redirect the user to /home
	context.Redirect(http.StatusSeeOther, "/api/v1/user/home")
}

// UserRegisterForm handles the HTTP GET request for the registration form.
//
// It checks for an existing Authorization cookie. If it exists, it redirects
// to /home. If not, it renders the register.html template with an empty
// gin.H object.
func UserRegisterForm(context *gin.Context) {
	// Use the helper function to check for redirection
	if middlewares.RedirectIfAuthenticated(context) {
		return
	}

	// If no token, render the register page
	context.HTML(http.StatusOK, "register.html", gin.H{})
}
