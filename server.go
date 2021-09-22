package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/JeanCntrs/admin-system/utils"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Main page")
	})

	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		friends := []string{"Marcela", "Laura", "Constanza", "Camila"}
		person := Person{PersonID: 1, Names: "Jean Carlos", Surnames: "Contreras Contreras", Age: 27, IsMan: true, Friends: friends}

		template := template.Must(template.ParseFiles("./html/product.html"))

		template.Execute(w, person)
	})

	http.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		person := Person{PersonID: 1, Names: "Jean Carlos", Surnames: "Contreras Contreras", Age: 27, IsMan: true}

		template, templateErr := template.ParseFiles("./html/categories.html")
		if templateErr != nil {
			panic("An error occurred when generating the categories template")
		}

		template.Execute(w, person)
	})

	http.HandleFunc("/not-found", func(w http.ResponseWriter, r *http.Request) {
		url := utils.GenerateURL("/products", "localhost:8000", "http", nil)
		response := utils.SendRequest("GET", url)
		fmt.Println("response request:", response)
		http.NotFound(w, r)
	})

	http.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Server error", 500)
	})

	http.ListenAndServe(":8000", nil)
}

type Person struct {
	PersonID int
	Names    string
	Surnames string
	Age      int
	IsMan    bool
	Friends  []string
}
