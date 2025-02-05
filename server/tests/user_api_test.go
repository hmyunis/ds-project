package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"context"
	"server/internal/user"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// MockUserService implements the user.Service interface for testing
type MockUserService struct{}

func (m *MockUserService) CreateUser(ctx context.Context, req *user.CreateUserReq) (*user.CreateUserRes, error) {
	return &user.CreateUserRes{ID: "123", Username: req.Username}, nil
}

func (m *MockUserService) Login(ctx context.Context, req *user.LoginUserReq) (*user.LoginUserRes, error) {
	return &user.LoginUserRes{}, nil
}

func TestUserSignup(t *testing.T) {
	router := gin.Default()

	mockService := &MockUserService{}       // ✅ Create the mock service
	handler := user.NewHandler(mockService) // ✅ Pass it to NewHandler
	router.POST("/signup", handler.CreateUser)

	signupData := `{"username": "testuser", "email": "test@example.com", "password": "securepass"}`
	req, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer([]byte(signupData)))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code, "Expected 201 Created")
}

func TestUserLogin(t *testing.T) {
	router := gin.Default()

	mockService := &MockUserService{}       // ✅ Create the mock service
	handler := user.NewHandler(mockService) // ✅ Pass it to NewHandler
	router.POST("/login", handler.Login)

	reqBody, _ := json.Marshal(map[string]string{
		"email":    "test@example.com",
		"password": "password123",
	})

	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
}
