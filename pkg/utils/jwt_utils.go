package utils

import (
	"time"

	"BuildasTechnicalAssessmentGo/pkg/models"
	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte("your_secret_key")

// GenerateJWT generates a JWT token for a given user
func GenerateJWT(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"sub":      user.ID,                               // User ID as a subject
		"exp":      time.Now().Add(24 * time.Hour).Unix(), // Token expiration
		"iat":      time.Now().Unix(),                     // Issued at
		"username": user.Username,                         // Additional user info
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
