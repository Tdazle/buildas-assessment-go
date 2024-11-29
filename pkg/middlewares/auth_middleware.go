package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
)

// AuthMiddleware is a middleware that checks the presence and validity of a JWT token
// in the Authorization cookie. It stores the parsed claims in the context for later
// use by other handlers_test. If the token is missing or invalid, it sends an error response
// with HTTP status 401 and aborts the request chain.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the Authorization header
		token, err := c.Cookie("Authorization")
		if token == "" {
			c.HTML(http.StatusUnauthorized, "error.html", gin.H{"error": "Authorization token not provided"})
			c.Abort()
			return
		}

		// Optionally, trim the "Bearer " prefix from the token
		token = strings.TrimPrefix(token, "Bearer ")

		// Parse the token and check validity
		claims, err := parseJWTToken(token)
		if err != nil {
			c.HTML(http.StatusUnauthorized, "error.html", gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Store the claims in the context for later use
		c.Set("claims", claims)

		// Proceed to the next handler
		c.Next()
	}
}

// parseJWTToken takes a JWT token and returns the claims if the token is valid
//
// It takes a string argument containing the JWT token, and returns a pointer
// to jwt.MapClaims and an error. If the token is invalid or missing, it
// returns an error. If the token is valid, it returns a pointer to the claims
// and nil.
func parseJWTToken(token string) (*jwt.MapClaims, error) {
	// Parse the JWT token and validate it
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Provide the secret key used for signing the token
		return []byte("your_secret_key"), nil
	})

	if err != nil || !parsedToken.Valid {
		return nil, err
	}

	// Get the claims from the parsed token
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return &claims, nil
}

// RedirectIfAuthenticated checks for the presence of an Authorization cookie.
//
// If the cookie is present and valid, it redirects the user to the "/api/v1/user/home" endpoint.
// It returns true if a redirection occurs and false otherwise.
func RedirectIfAuthenticated(context *gin.Context) bool {
	token, err := context.Cookie("Authorization")
	if err == nil && token != "" {
		context.Redirect(http.StatusSeeOther, "/api/v1/user/home")
		return true
	}
	return false
}
