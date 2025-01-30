package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {
	// edit with your own mysql info
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/chatapp")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("error verifying connection with db.Ping")
		panic(err.Error())
	}

	return &Database{db: db}, nil
}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}
