package repository

import (
	"BuildasTechnicalAssessmentGo/pkg/models"
	"errors"
	"gorm.io/gorm"
)

// PostgresUserRepository implements UserRepository interface for PostgresSQL
type PostgresUserRepository struct {
	DB *gorm.DB
}

// CreateUser creates a new user in the database
func (r *PostgresUserRepository) CreateUser(user *models.User) error {
	return r.DB.Create(user).Error
}

// GetUserByUsername retrieves a user by their username
func (r *PostgresUserRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	result := r.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil // User not found
		}
		return nil, result.Error // Some other error occurred
	}
	return &user, nil // User found
}

// GetAllUsers retrieves all users
func (r *PostgresUserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := r.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
