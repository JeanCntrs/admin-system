package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/JeanCntrs/admin-system/dal"
	"github.com/JeanCntrs/admin-system/utils"
	"github.com/gorilla/mux"
)

func Provider(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "provider", nil)
}

func GetProviders(w http.ResponseWriter, r *http.Request) {
	providers := dal.GetProviders()
	providersByte, _ := json.Marshal(providers)
	fmt.Fprint(w, string(providersByte))
}

func GetProvidersByCountryId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	countryId := vars["countryId"]
	countryIdConv, _ := strconv.Atoi(countryId)

	providers := dal.GetProvidersByCountryId(countryIdConv)
	providersByte, _ := json.Marshal(providers)
	fmt.Fprint(w, string(providersByte))
}
