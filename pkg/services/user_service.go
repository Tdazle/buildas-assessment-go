package services

import (
	"BuildasTechnicalAssessmentGo/internal/repository"
	"BuildasTechnicalAssessmentGo/pkg/models"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

// UserServiceInterface defines the interface for the UserService
type UserServiceInterface interface {
	RegisterUser(username, password string) error
	GetUserByUsername(username string) (*models.User, error)
	CheckPassword(hashedPassword, plainPassword string) error
	GetAllUsers() ([]models.User, error)
}

// UserService contains methods for managing users
type UserService struct {
	Repo repository.UserRepository
}

// NewUserService creates a new UserService
func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

// RegisterUser handles user registration
func (s *UserService) RegisterUser(username, password string) error {
	// Check if the user already exists using the repository
	existingUser, err := s.Repo.GetUserByUsername(username)
	if err != nil {
		return err // Return the error if the query fails
	}
	if existingUser != nil {
		return errors.New("user already exists") // Return specific error
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Create new user
	newUser := models.User{
		Username: username,
		Password: string(hashedPassword),
	}

	// Save to the database using the repository
	err = s.Repo.CreateUser(&newUser)
	if err != nil {
		return err
	}

	return nil
}

// GetUserByUsername uses the repository to fetch a user by their username
func (s *UserService) GetUserByUsername(username string) (*models.User, error) {
	return s.Repo.GetUserByUsername(username)
}

// CheckPassword compares a plain-text password with a hashed password
func (s *UserService) CheckPassword(hashedPassword, plainPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err != nil {
		return errors.New("invalid password")
	}
	return nil
}

// GetAllUsers uses the repository to fetch all users
func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.Repo.GetAllUsers()
}
