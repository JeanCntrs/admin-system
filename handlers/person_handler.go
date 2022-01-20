package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/JeanCntrs/admin-system/dal"
	"github.com/JeanCntrs/admin-system/models"
	"github.com/JeanCntrs/admin-system/utils"
	"github.com/gorilla/mux"
)

func Person(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "person", nil)
}

func GetPersons(w http.ResponseWriter, r *http.Request) {
	persons := dal.GetPersons()
	personsByte, _ := json.Marshal(persons)
	fmt.Fprint(w, string(personsByte))
}

func GetPersonById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	PersonId := vars["personId"]
	personIdConv, _ := strconv.Atoi(PersonId)

	persons := dal.GetPersonById(personIdConv)
	personsByte, _ := json.Marshal(persons)
	fmt.Fprint(w, string(personsByte))
}

func GetTypePersons(w http.ResponseWriter, r *http.Request) {
	typePersons := dal.GetTypePersons()
	typePersonsByte, _ := json.Marshal(typePersons)
	fmt.Fprint(w, string(typePersonsByte))
}

func GetPersonsByFullname(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fullname := vars["fullname"]

	persons := dal.GetPersonsByFullname(fullname)
	personsByte, _ := json.Marshal(persons)
	fmt.Fprint(w, string(personsByte))
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	person := models.Person{}
	data := json.NewDecoder(r.Body)
	err := data.Decode(&person)

	if err != nil {
		panic("An error occurred while decoding country")
	}

	if person.PersonId == 0 {
		_, err := dal.InsertPerson(person)
		if err != nil {
			fmt.Fprintf(w, "0")
		}
	} else {
		_, err := dal.UpdatePerson(person)
		if err != nil {
			fmt.Fprintf(w, "0")
		}
	}

	fmt.Fprintf(w, "1")
}
