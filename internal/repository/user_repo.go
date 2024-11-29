package repository

import "BuildasTechnicalAssessmentGo/pkg/models"

// UserRepository defines methods for interacting with users in the database
type UserRepository interface {
	// CreateUser adds a new user to the database.
	//
	// It takes a pointer to a User struct as an argument and returns an error.
	// If the user is successfully created, it will return nil.
	CreateUser(user *models.User) error

	// GetUserByUsername fetches a user by their username from the database.
	//
	// It takes a string argument representing the username and returns a pointer
	// to a User struct and an error. If the user is found, it will return the
	// user and nil. If the user is not found, it will return nil and an error.
	GetUserByUsername(username string) (*models.User, error)

	// GetAllUsers retrieves all users from the database.
	//
	// It returns a slice of User structs and an error. If the users are
	// successfully retrieved, it will return the slice and nil. If there
	// is an error, it will return nil and an error.
	GetAllUsers() ([]models.User, error)
}
