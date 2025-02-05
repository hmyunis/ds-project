package tests

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestMySQLConnection(t *testing.T) {
	dsn := "root:2405@tcp(localhost:3306)/chatapp"
	db, err := sql.Open("mysql", dsn)
	assert.NoError(t, err, "Database connection failed")

	defer db.Close()
	err = db.Ping()
	assert.NoError(t, err, "Database ping failed")
}
