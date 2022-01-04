package main

import (
	"fmt"
	"net/http"

	"github.com/JeanCntrs/admin-system/dal"
	"github.com/JeanCntrs/admin-system/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// files := http.FileServer(http.Dir("static"))
	// http.Handle("/static/", http.StripPrefix("/static/", files))

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	r.HandleFunc("/products", handlers.Product)
	r.HandleFunc("/products/create", handlers.CreateProduct)
	r.HandleFunc("/products/save", handlers.SaveProduct)
	r.HandleFunc("/products/edit/{id}", handlers.EditProduct)
	r.HandleFunc("/products/delete/{id}", handlers.DeleteProduct)

	r.HandleFunc("/categories", handlers.Category)
	r.HandleFunc("/categories/create", handlers.CreateCategory)
	r.HandleFunc("/categories/edit/{id}", handlers.EditCategory)
	r.HandleFunc("/categories/delete/{id}", handlers.DeleteCategory)

	r.HandleFunc("/countries", handlers.Country)
	r.HandleFunc("/countries/list", handlers.GetCountries)

	r.HandleFunc("/providers", handlers.Provider)

	r.HandleFunc("/not-found", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	r.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Server error", 500)
	})

	r.HandleFunc("/list-categories", func(w http.ResponseWriter, r *http.Request) {
		categoryList := dal.ListCategories()

		for i, v := range categoryList {
			fmt.Println("i", i)
			fmt.Println("v", v.Name)
		}
	})

	fmt.Println("Server started on port :8000")

	http.ListenAndServe(":8000", r)
}
