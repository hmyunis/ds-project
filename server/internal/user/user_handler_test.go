package user

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockService is a mock implementation of the Service interface
type MockService struct {
	mock.Mock
}

func (m *MockService) CreateUser(c context.Context, req *CreateUserReq) (*CreateUserRes, error) {
	args := m.Called(c, req)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*CreateUserRes), nil
}

func (m *MockService) Login(c context.Context, req *LoginUserReq) (*LoginUserRes, error) {
	args := m.Called(c, req)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*LoginUserRes), nil
}

func TestCreateUserHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Arrange
	mockService := new(MockService)
	handler := NewHandler(mockService)

	createUserReq := &CreateUserReq{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password",
	}

	expectedResponse := &CreateUserRes{
		ID:       "1",
		Username: "testuser",
		Email:    "test@example.com",
	}

	mockService.On("CreateUser", mock.Anything, createUserReq).Return(expectedResponse, nil)

	// Create a Gin context and request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	jsonValue, _ := json.Marshal(createUserReq)
	req, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	// Act
	handler.CreateUser(c)

	// Assert
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), `"username":"testuser"`)
	mockService.AssertExpectations(t) // Verify that the mock service was called as expected
}

func TestLoginHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Arrange
	mockService := new(MockService)
	handler := NewHandler(mockService)

	loginUserReq := &LoginUserReq{
		Email:    "test@example.com",
		Password: "password",
	}

	expectedResponse := &LoginUserRes{
		ID:       "1",
		Username: "testuser",
	}

	mockService.On("Login", mock.Anything, loginUserReq).Return(expectedResponse, nil)

	// Create a Gin context and request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	jsonValue, _ := json.Marshal(loginUserReq)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	// Act
	handler.Login(c)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"username":"testuser"`)
	mockService.AssertExpectations(t) // Verify that the mock service was called as expected
}

func TestLogoutHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Arrange
	mockService := new(MockService)
	handler := NewHandler(mockService)

	// Create a Gin context and request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest("GET", "/logout", nil)
	c.Request = req

	// Act
	handler.Logout(c)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"message":"logout successful"`)
	assert.NotEmpty(t, c.Writer.Header().Get("Set-Cookie"))
}
