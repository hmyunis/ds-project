package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

// MockUserHandler and MockWSHandler are mock implementations
// You would add the appropriate mocked methods as needed for your routes.

type MockUserHandler struct {
	mock.Mock
}

type MockWSHandler struct {
	mock.Mock
}

func (m *MockUserHandler) CreateUser(c *gin.Context) {
	m.Called(c)
	c.String(http.StatusOK, "CreateUser called") // dummy response
}

func (m *MockUserHandler) Login(c *gin.Context) {
	m.Called(c)
	c.String(http.StatusOK, "Login called") // dummy response
}

func (m *MockUserHandler) Logout(c *gin.Context) {
	m.Called(c)
	c.String(http.StatusOK, "Logout called") // dummy response
}
