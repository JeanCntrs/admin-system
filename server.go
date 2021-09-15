package main

import (
	"fmt"
	"net/http"
	"net/url"
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

		uri := "/carts"
		host := "localhost:8000"
		protocol := "http"
		urlParams := map[string]string{"id": "1", "name": "cart_1"}

		generatedURL := generateURL(uri, host, protocol, urlParams)
		fmt.Println("generatedURL", generatedURL)
	})

	http.HandleFunc("/not-found", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	http.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Server error", 500)
	})

	http.ListenAndServe(":8000", nil)
}

func generateURL(uri, host, protocol string, urlParams map[string]string) string {
	url, _ := url.Parse(uri)
	url.Host = host
	url.Scheme = protocol
	mapFunction := url.Query()

	for key, value := range urlParams {
		mapFunction.Add(key, value)
	}

	url.RawQuery = mapFunction.Encode()
	return url.String()

}
