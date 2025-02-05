package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	// Test hashing a valid password
	password := "testpassword"
	hashedPassword, err := HashPassword(password)
	assert.NoError(t, err, "HashPassword should not return an error")
	assert.NotEmpty(t, hashedPassword, "Hashed password should not be empty")

	// Test hashing an empty password
	emptyPassword := ""
	hashedEmptyPassword, err := HashPassword(emptyPassword)
	assert.NoError(t, err, "HashPassword should not return an error for an empty password")
	assert.NotEmpty(t, hashedEmptyPassword, "Hashed empty password should not be empty")
}

func TestCheckPassword(t *testing.T) {
	// Test checking a valid password
	password := "testpassword"
	hashedPassword, err := HashPassword(password)
	assert.NoError(t, err, "HashPassword should not return an error")

	err = CheckPassword(password, hashedPassword)
	assert.NoError(t, err, "CheckPassword should not return an error for a valid password")

	// Test checking an incorrect password
	incorrectPassword := "wrongpassword"
	err = CheckPassword(incorrectPassword, hashedPassword)
	assert.Error(t, err, "CheckPassword should return an error for an incorrect password")

	// Test checking an empty password
	emptyPassword := ""
	hashedEmptyPassword, err := HashPassword(emptyPassword)
	assert.NoError(t, err, "HashPassword should not return an error for an empty password")

	err = CheckPassword(emptyPassword, hashedEmptyPassword)
	assert.NoError(t, err, "CheckPassword should not return an error for an empty password")

	// Test checking a password against an invalid hash
	invalidHash := "invalidhash"
	err = CheckPassword(password, invalidHash)
	assert.Error(t, err, "CheckPassword should return an error for an invalid hash")
}
