package database

import (
	"database/sql"
	"errors"
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
		return nil, errors.New("error executing the query")
	}

	return rows, err
}

func Excec(query string, args ...interface{}) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	if err != nil {
		return nil, errors.New("error inserting, updating or deleting")
	}

	return result, err
}

func QueryRow(query string, args ...interface{}) (int, error) {
	var quantity int
	err := db.QueryRow(query, args...).Scan(&quantity)
	if err != nil {
		return -1, err
	}

	return quantity, err
}

func Begin() (*sql.Tx, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func CloseConnection() {
	db.Close()
}
