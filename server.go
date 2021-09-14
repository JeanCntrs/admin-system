package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Main page")
	})

	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Products page")

		mapURL := r.URL.Query()
		fmt.Println("mapURL", mapURL)

		mapURL.Add("Price", "10")
		fmt.Println("mapURL.Encode()", mapURL.Encode())

		fmt.Println("URL", r.URL)
		fmt.Println("RawQuery", r.URL.RawQuery)
		fmt.Println("Name", r.URL.Query().Get("name"))
		fmt.Println("ProductId", r.URL.Query().Get("productId"))
	})

	http.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Categories page")
		// Redirect to main page
		// http.Redirect(w, r, "/", 301)
	})

	http.HandleFunc("/not-found", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	http.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Server error", 500)
	})

	http.ListenAndServe(":8000", nil)
}
