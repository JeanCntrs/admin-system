package handlers

import (
	"net/http"

	"github.com/JeanCntrs/admin-system/models"
	"github.com/JeanCntrs/admin-system/utils"
)

func Category(w http.ResponseWriter, r *http.Request) {
	person := models.Person{PersonID: 1, Names: "Jean Carlos", Surnames: "Contreras Contreras", Age: 27, IsMan: true}

	utils.RenderTemplate(w, "category", person)
}
