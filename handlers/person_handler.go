package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JeanCntrs/admin-system/dal"
	"github.com/JeanCntrs/admin-system/utils"
)

func Person(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "person", nil)
}

func GetPersons(w http.ResponseWriter, r *http.Request) {
	persons := dal.GetPersons()
	personsByte, _ := json.Marshal(persons)
	fmt.Fprint(w, string(personsByte))
}
