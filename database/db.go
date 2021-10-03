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

type Category struct {
	idcategoria int
	nombre      string
	descripcion string
}

type CategoryList []Category

func ConnectDB() {
	connStr := "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic("Connection error")
	}

	defer db.Close()

	rows, err := db.Query(`SELECT idcategoria, nombre, descripcion FROM public.categoria`)
	if err != nil {
		fmt.Println("errorrrr", err)
		panic("Error executing the query")
	}

	categoryList := CategoryList{}
	for rows.Next() {
		category := Category{}
		rows.Scan(&category.idcategoria, &category.nombre, &category.descripcion)
		categoryList = append(categoryList, category)
	}

	for i, v := range categoryList {
		fmt.Println("i", i)
		fmt.Println("v", v.nombre)
	}

	fmt.Println("Successful connection")
}
