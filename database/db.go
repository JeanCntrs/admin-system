package database

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"

	"github.com/JeanCntrs/admin-system/config"
)

var db *sql.DB

func OpenConnection() {
	config.LoadEnv()

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	connStr := "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=disable"

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic("Connection error")
	}
}

func Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		panic("Error executing the query")
	}

	return rows, err
}

func CloseConnection() {
	db.Close()
}
