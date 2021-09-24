package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/JeanCntrs/admin-system/utils"
)

var funcsMap = template.FuncMap{"Welcome": Welcome}

var allTemplates = template.Must(template.New("T").Funcs(funcsMap).ParseFiles(
	"./html/person/person.html",
	"./html/main/index.html",
	"./html/category/category.html",
	"./html/product/product.html",
	"./html/includes/message.html",
))

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Main page")
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

		allTemplates.ExecuteTemplate(w, "category", person)
	})

	http.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		person := Person{PersonID: 1, Names: "Jean Carlos", Surnames: "Contreras Contreras", Age: 27, IsMan: true}

		template, templateErr := template.ParseFiles("./html/category/category.html")
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
