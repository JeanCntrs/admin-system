package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = ""
	dbname   = "dbsupermarket"
)

var db *sql.DB

func OpenConnection() {
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
