package handlers

import (
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(username, password string) error {
	// Check if the user already exists
	var exists bool
	err := DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)", username).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return errors.New("database query error")
	}
	if exists {
		return errors.New("username already exists")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}

	// Insert the new user into the database
	_, err = DB.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, string(hashedPassword))
	if err != nil {
		return errors.New("failed to register user")
	}

	return nil
}

func AuthenticateUser(username, password string) error {
	// Retrieve the hashed password from the database
	var hashedPassword string
	err := DB.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&hashedPassword)
	if err == sql.ErrNoRows {
		return errors.New("invalid username or password")
	} else if err != nil {
		return errors.New("database query error")
	}

	// Compare the provided password with the hashed password
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return errors.New("invalid username or password")
	}

	return nil
}
