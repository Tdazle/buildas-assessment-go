package services_mock

import (
	"BuildasTechnicalAssessmentGo/pkg/models"
	"BuildasTechnicalAssessmentGo/pkg/services"
	"github.com/stretchr/testify/mock"
)

// MockUserService should implement the UserService interface
type MockUserService struct {
	mock.Mock
}

// Ensure that MockUserService implements UserServiceInterface
var _ services.UserServiceInterface = (*MockUserService)(nil)

// GetUserByUsername Implement the methods of UserService
func (m *MockUserService) GetUserByUsername(username string) (*models.User, error) {
	args := m.Called(username)
	if user := args.Get(0); user != nil {
		return user.(*models.User), args.Error(1)
	}
	return nil, args.Error(1)
}

// RegisterUser Implement the methods of UserService
func (m *MockUserService) RegisterUser(username, password string) error {
	args := m.Called(username, password)
	return args.Error(0)
}

// GetAllUsers Implement the methods of UserService
func (m *MockUserService) GetAllUsers() ([]models.User, error) {
	args := m.Called()
	if users := args.Get(0); users != nil {
		return users.([]models.User), args.Error(1)
	}
	return nil, args.Error(1)
}

// CheckPassword Implement the methods of UserService
func (m *MockUserService) CheckPassword(hashedPassword, password string) error {
	args := m.Called(hashedPassword, password)
	return args.Error(0)
}
