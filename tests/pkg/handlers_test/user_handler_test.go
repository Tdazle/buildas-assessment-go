package handlers_test

import (
	"BuildasTechnicalAssessmentGo/internal/repository/repository_mock"
	"BuildasTechnicalAssessmentGo/pkg/handlers"
	"BuildasTechnicalAssessmentGo/pkg/models"
	"BuildasTechnicalAssessmentGo/pkg/services"
	"BuildasTechnicalAssessmentGo/pkg/services/services_mock"
	"bytes"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRegisterUser(t *testing.T) {
	tests := []struct {
		name         string
		requestBody  string
		mockBehavior func(mock *repository_mock.MockUserRepository)
		expectedCode int
	}{
		{
			name:        "Valid Registration",
			requestBody: `{"username":"newuser","password":"password123"}`,
			mockBehavior: func(mock *repository_mock.MockUserRepository) {
				mock.On("GetUserByUsername", "newuser").Return(nil, nil) // Simulate no existing user
				mock.On("CreateUser", &models.User{Username: "newuser", Password: "password123"}).Return(nil)
			},
			expectedCode: http.StatusOK,
		},
		{
			name:        "Duplicate User",
			requestBody: `{"username":"existinguser","password":"password123"}`,
			mockBehavior: func(mock *repository_mock.MockUserRepository) {
				mock.On("CreateUser", &models.User{Username: "existinguser", Password: "password123"}).Return(errors.New("user already exists"))
			},
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			c.Request = httptest.NewRequest("POST", "/register", bytes.NewBufferString(tt.requestBody))
			c.Request.Header.Set("Content-Type", "application/json")

			// Create mock repository instance
			mockRepo := new(repository_mock.MockUserRepository)
			tt.mockBehavior(mockRepo)

			// Pass the mock repository to the userService constructor
			userService := services.NewUserService(mockRepo)

			// Call the handler with the mock service
			handlers.RegisterUser(c, userService)

			// Check if the status code is as expected
			if w.Code != tt.expectedCode {
				t.Errorf("Expected status %d but got %d", tt.expectedCode, w.Code)
			}

			// Verify that the mock expectations were met
			mockRepo.AssertExpectations(t)
		})

	}
}

func TestAddUser(t *testing.T) {
	router := gin.Default()
	mockService := new(services_mock.MockUserService)

	// Mock the service response
	mockService.On("RegisterUser", "NewUser", "password123").Return(nil)

	router.POST("/api/v1/user/add", func(c *gin.Context) {
		handlers.AddUser(c, mockService)
	})

	form := url.Values{}
	form.Add("username", "NewUser")
	form.Add("password", "password123")
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/user/add", strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusSeeOther, w.Code)
	assert.Equal(t, "/api/v1/user/home", w.Header().Get("Location"))
}

func TestLoginUser(t *testing.T) {
	router := gin.Default()
	mockService := new(services_mock.MockUserService)

	user := &models.User{Username: "TestUser", Password: "hashedpassword"}
	mockService.On("GetUserByUsername", "TestUser").Return(user, nil)
	mockService.On("CheckPassword", "hashedpassword", "password123").Return(nil)

	router.POST("/api/v1/user/login", func(c *gin.Context) {
		handlers.LoginUser(c, mockService)
	})

	form := url.Values{}
	form.Add("username", "TestUser")
	form.Add("password", "password123")
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/user/login", strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "TestUser")
}

func TestRegisterUser2(t *testing.T) {
	router := gin.Default()
	mockService := new(services_mock.MockUserService)

	mockService.On("RegisterUser", "NewUser", "password123").Return(nil)
	mockService.On("GetUserByUsername", "NewUser").Return(&models.User{Username: "NewUser"}, nil)

	router.POST("/api/v1/user/register", func(c *gin.Context) {
		handlers.RegisterUser(c, mockService)
	})

	form := url.Values{}
	form.Add("username", "NewUser")
	form.Add("password", "password123")
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/user/register", strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusSeeOther, w.Code)
	assert.Equal(t, "/api/v1/user/home", w.Header().Get("Location"))
}

func TestLoginUserForm(t *testing.T) {
	router := gin.Default()
	router.GET("/api/v1/user/login/form", handlers.LoginUserForm)

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/user/login/form", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "login.html")
}

func TestUserRegisterForm(t *testing.T) {
	router := gin.Default()
	router.GET("/api/v1/user/register", handlers.UserRegisterForm)

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/user/register", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "register.html")
}

func TestHome(t *testing.T) {
	router := gin.Default()
	mockService := new(services_mock.MockUserService)

	// Mock the service response
	mockService.On("GetAllUsers").Return([]models.User{
		{Username: "User1"},
		{Username: "User2"},
	}, nil)

	// Add claims to context
	router.GET("/api/v1/user/home", func(c *gin.Context) {
		claims := &jwt.MapClaims{"username": "TestUser"}
		c.Set("claims", claims)
		handlers.Home(c, mockService)
	})

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/user/home", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "TestUser")
	assert.Contains(t, w.Body.String(), "User1")
	assert.Contains(t, w.Body.String(), "User2")
}
