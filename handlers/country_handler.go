package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JeanCntrs/admin-system/dal"
	"github.com/JeanCntrs/admin-system/utils"
	"github.com/gorilla/mux"
)

func Country(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "country", nil)
}

func GetCountries(w http.ResponseWriter, r *http.Request) {
	countries := dal.GetCountries()
	countriesByte, _ := json.Marshal(countries)
	fmt.Fprint(w, string(countriesByte))
}

func GetCountriesByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	countries := dal.GetCountriesByName(name)
	countriesByte, _ := json.Marshal(countries)
	fmt.Fprint(w, string(countriesByte))
}
