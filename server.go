package main

import (
	"net/http"

	"github.com/JeanCntrs/admin-system/database"
	"github.com/JeanCntrs/admin-system/handlers"
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
		database.ConnectDB()
	})

	http.ListenAndServe(":8000", nil)
}
