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

// Setting up the MockUserService that implements the eser.service interface 

type MockUserService struct{}

func (m *MockUserService) CreateUser(ctx context.Context, req *user.CreateUserReq) (*user.CreateUserRes, error) {
	return &user.CreateUserRes{ID: "123", Username: req.Username}, nil
}

func (m *MockUserService) Login(ctx context.Context, req *user.LoginUserReq) (*user.LoginUserRes, error) {
	return &user.LoginUserRes{}, nil
}
//test the signup functionality(test if a user is successfully created)
func TestUserSignup(t *testing.T) {
	router := gin.Default()

	mockService := &MockUserService{}       
	handler := user.NewHandler(mockService) 
	router.POST("/signup", handler.CreateUser)

	signupData := `{"username": "testuser", "email": "test@example.com", "password": "securepass"}`
	req, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer([]byte(signupData)))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code, "Expected 201 Created")
}

// test the login functionality
func TestUserLogin(t *testing.T) {
	router := gin.Default()

	mockService := &MockUserService{}       
	handler := user.NewHandler(mockService) 
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
