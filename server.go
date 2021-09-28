package main

import (
	"fmt"
	"net/http"

	"github.com/JeanCntrs/admin-system/utils"
)

func main() {
	files := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		utils.RenderTemplate(w, "index", nil)
	})

	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		hobbies := []Hobby{
			{Name: "Run", Description: "Excercise"},
			{Name: "Study", Description: "Online courses"},
		}
		friends := []string{"Marcela", "Laura", "Constanza", "Camila"}
		person := Person{
			PersonID: 1,
			Names:    "Jean Carlos",
			Surnames: "Contreras Contreras",
			Age:      27,
			IsMan:    true,
			Friends:  friends,
			Hobbies:  hobbies,
		}

		utils.RenderTemplate(w, "product", person)
	})

	http.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		person := Person{PersonID: 1, Names: "Jean Carlos", Surnames: "Contreras Contreras", Age: 27, IsMan: true}

		utils.RenderTemplate(w, "category", person)
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

type Hobby struct {
	Name        string
	Description string
}

type Person struct {
	PersonID int
	Names    string
	Surnames string
	Age      int
	IsMan    bool
	Friends  []string
	Hobbies  []Hobby
}
