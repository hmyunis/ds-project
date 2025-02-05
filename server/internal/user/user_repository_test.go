package user

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestRepository_CreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewRepository(db)

	user := &User{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "testpass",
	}

	mock.ExpectExec("INSERT INTO users").
		WithArgs(user.Username, user.Password, user.Email).
		WillReturnResult(sqlmock.NewResult(1, 1))

	res, err := repo.CreateUser(context.Background(), user)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), res.ID)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRepository_GetUserByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewRepository(db)

	rows := sqlmock.NewRows([]string{"id", "email", "username", "password"}).
		AddRow(1, "test@example.com", "testuser", "testpass")

	mock.ExpectQuery("SELECT id, email, username, password FROM users WHERE email = ?").
		WithArgs("test@example.com").
		WillReturnRows(rows)

	user, err := repo.GetUserByEmail(context.Background(), "test@example.com")
	assert.NoError(t, err)
	assert.Equal(t, "testuser", user.Username)
	assert.NoError(t, mock.ExpectationsWereMet())
}
