package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = ""
	dbname   = "dbsupermarket"
)

func ConnectDB() {
	connStr := "host=" + host + "port=" + port + "user=" + user + "password=" + password + "dbname=" + dbname + "sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic("Connection error")
	}

	defer db.Close()
	fmt.Println("Successful connection")
}
