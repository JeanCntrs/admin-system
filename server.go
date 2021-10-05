package main

import (
	"fmt"
	"net/http"

	"github.com/JeanCntrs/admin-system/database"
	"github.com/JeanCntrs/admin-system/handlers"
	"github.com/JeanCntrs/admin-system/models"
)

func main() {
	files := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	http.HandleFunc("/", handlers.Index)

	http.HandleFunc("/products", handlers.Product)

	http.HandleFunc("/categories", handlers.Category)

	http.HandleFunc("/persons", handlers.Person)

	http.HandleFunc("/not-found", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	http.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Server error", 500)
	})

	http.HandleFunc("/conn", func(w http.ResponseWriter, r *http.Request) {
		query := `SELECT idcategoria, nombre, descripcion FROM public.categoria`
		database.OpenConnection()
		rows, _ := database.Query(query)
		database.CloseConnection()

		type CategoryList []models.Category

		categoryList := CategoryList{}
		for rows.Next() {
			category := models.Category{}
			rows.Scan(&category.Idcategoria, &category.Nombre, &category.Descripcion)
			categoryList = append(categoryList, category)
		}

		for i, v := range categoryList {
			fmt.Println("i", i)
			fmt.Println("v", v.Nombre)
		}
	})

	http.ListenAndServe(":8000", nil)
}
