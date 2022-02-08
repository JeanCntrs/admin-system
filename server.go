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
	r.HandleFunc("/countries/create", handlers.CreateCountry)
	r.HandleFunc("/countries/list", handlers.GetCountries)
	r.HandleFunc("/countries/list/name/{name}", handlers.GetCountriesByName)
	r.HandleFunc("/countries/list/id/{countryId}", handlers.GetCountryById)
	r.HandleFunc("/countries/delete/{id}", handlers.DeleteCountry)

	r.HandleFunc("/providers", handlers.Provider)
	r.HandleFunc("/providers/create", handlers.CreateProvider)
	r.HandleFunc("/providers/list", handlers.GetProviders)
	r.HandleFunc("/providers/list/id/{countryId}", handlers.GetProvidersByCountryId)
	r.HandleFunc("/providers/{providerId}", handlers.GetProviderById)
	r.HandleFunc("/providers/delete/{id}", handlers.DeleteProvider)

	r.HandleFunc("/persons", handlers.Person)
	r.HandleFunc("/persons/create", handlers.CreatePerson)
	r.HandleFunc("/persons/list", handlers.GetPersons)
	r.HandleFunc("/persons/list/without-user", handlers.GetPersonsWithoutUser)
	r.HandleFunc("/persons/list/type", handlers.GetTypePersons)
	r.HandleFunc("/persons/list/name/{fullname}", handlers.GetPersonsByFullname)
	r.HandleFunc("/persons/list/id/{personId}", handlers.GetPersonById)
	r.HandleFunc("/persons/delete/{id}", handlers.DeletePerson)

	r.HandleFunc("/users", handlers.User)
	r.HandleFunc("/users/list", handlers.GetUsers)
	r.HandleFunc("/users/list/id/{userId}", handlers.GetUserById)
	r.HandleFunc("/users/create", handlers.CreateUser)
	r.HandleFunc("/users/delete/{id}", handlers.DeleteUser)

	r.HandleFunc("/pages", handlers.Page)
	r.HandleFunc("/pages/list", handlers.GetPages)
	r.HandleFunc("/pages/list/id/{pageId}", handlers.GetPageById)
	r.HandleFunc("/pages/create", handlers.CreatePage)

	r.HandleFunc("/role-page", handlers.RolePage)
	r.HandleFunc("/role-page/create", handlers.CreateRolePage)

	r.HandleFunc("/role-types/create", handlers.CreateRoleType)
	r.HandleFunc("/role-types/list", handlers.GetRoleTypes)

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
