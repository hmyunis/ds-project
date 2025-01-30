package handlers

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// InitializeDatabase initializes the MySQL database connection
func InitializeDatabase() {
	var err error

	// Replace with your MySQL credentials
	DB, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/chatapp")
	if err != nil {
		log.Fatal("Failed to connect to MySQL:", err)
	}

	// Test the database connection
	err = DB.Ping()
	if err != nil {
		log.Fatal("Failed to ping MySQL:", err)
	}

	log.Println("Connected to MySQL database.")
}
