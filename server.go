package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/JeanCntrs/admin-system/utils"
)

var funcsMap = template.FuncMap{"Welcome": Welcome}

// var allTemplates = template.Must(template.New("T").Funcs(funcsMap).ParseFiles(
// 	"./html/category/category.html",
// 	"./html/includes/message.html",
// 	"./html/main/index.html",
// 	"./html/person/person.html",
// 	"./html/product/product.html",
// ))

var allTemplates = template.Must(template.New("T").Funcs(funcsMap).ParseGlob("./html/**/*.html"))
var errTemplate = template.Must(template.ParseFiles("./html/error/error.html"))

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

		err := allTemplates.ExecuteTemplate(w, "index", nil)
		if err != nil {
			w.WriteHeader(500)
			errTemplate.Execute(w, nil)
		}
	})

	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

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

		err := allTemplates.ExecuteTemplate(w, "product", person)
		if err != nil {
			w.WriteHeader(500)
			errTemplate.Execute(w, nil)
		}
	})

	http.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

		person := Person{PersonID: 1, Names: "Jean Carlos", Surnames: "Contreras Contreras", Age: 27, IsMan: true}

		err := allTemplates.ExecuteTemplate(w, "category", person)
		if err != nil {
			w.WriteHeader(500)
			errTemplate.Execute(w, nil)
		}
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

func (p Person) Greeting(name string) string {
	return "Hi" + p.Names
}

func Welcome(name string) string {
	return "Welcome to the page " + name
}
