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

func GetCountryById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	countryId := vars["countryId"]
	countryIdConv, _ := strconv.Atoi(countryId)

	countries := dal.GetCountryById(countryIdConv)
	countriesByte, _ := json.Marshal(countries)
	fmt.Fprint(w, string(countriesByte))
}

func CreateCountry(w http.ResponseWriter, r *http.Request) {
	country := models.Country{}
	data := json.NewDecoder(r.Body)
	err := data.Decode(&country)

	if err != nil {
		panic("An error occurred while decoding country")
	}

	if country.CountryId == 0 {
		_, err := dal.InsertCountry(country)
		if err != nil {
			fmt.Fprintf(w, "0")
		}
	} else {
		_, err := dal.UpdateCountry(country)
		if err != nil {
			fmt.Fprintf(w, "0")
		}
	}

	fmt.Fprintf(w, "1")
}

func DeleteCountry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	idConv, err := strconv.Atoi(id)

	if err != nil {
		fmt.Fprintf(w, "0")
		return
	}

	_, errorFound := dal.DeleteCountry(idConv)

	if errorFound != nil {
		fmt.Fprintf(w, "0")
		return
	}

	fmt.Fprintf(w, "1")
}
