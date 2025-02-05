package user

import (
	"context"
	"fmt"
	"server/util"
	"strconv"
	"testing"

	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) CreateUser(ctx context.Context, user *User) (*User, error) {
	args := m.Called(ctx, user)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*User), nil
}

func (m *MockRepository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	args := m.Called(ctx, email)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*User), nil
}

func TestCreateUserService(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	req := &CreateUserReq{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password",
	}

	hashedPassword, _ := util.HashPassword(req.Password)

	expectedUser := &User{
		ID:       1,
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	mockRepo.On("CreateUser", mock.Anything, mock.MatchedBy(func(user *User) bool {
		return user.Username == req.Username && user.Email == req.Email
	})).Return(expectedUser, nil)

	res, err := service.CreateUser(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, strconv.Itoa(int(expectedUser.ID)), res.ID)
	mockRepo.AssertExpectations(t)
}

func TestLogin(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	req := &LoginUserReq{
		Email:    "test@example.com",
		Password: "password",
	}

	hashedPassword, _ := util.HashPassword(req.Password)

	expectedUser := &User{
		ID:       1,
		Username: "testuser",
		Email:    req.Email,
		Password: hashedPassword,
	}

	mockRepo.On("GetUserByEmail", mock.Anything, req.Email).Return(expectedUser, nil)

	res, err := service.Login(context.Background(), req)

	assert.NoError(t, err)
	assert.NotEmpty(t, res.accessToken)
	assert.Equal(t, strconv.Itoa(int(expectedUser.ID)), res.ID)
	mockRepo.AssertExpectations(t)

	//Verify JWT
	token, _ := jwt.Parse(res.accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secretKey), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		assert.Equal(t, expectedUser.Username, claims["username"])
		assert.Equal(t, strconv.Itoa(int(expectedUser.ID)), claims["id"])
	} else {
		t.Error("Invalid Token")
	}

}
