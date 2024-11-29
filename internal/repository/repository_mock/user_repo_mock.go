package repository_mock

import (
	"BuildasTechnicalAssessmentGo/internal/repository"
	"BuildasTechnicalAssessmentGo/pkg/models"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

// Ensure MockUserRepository implements the UserRepository interface
var _ repository.UserRepository = (*MockUserRepository)(nil)

func (m *MockUserRepository) CreateUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) GetUserByUsername(username string) (*models.User, error) {
	args := m.Called(username)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) GetAllUsers() ([]models.User, error) {
	args := m.Called()
	return args.Get(0).([]models.User), args.Error(1)
}
