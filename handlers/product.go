package handlers

import (
	"net/http"

	"github.com/JeanCntrs/admin-system/models"
	"github.com/JeanCntrs/admin-system/utils"
)

func Product(w http.ResponseWriter, r *http.Request) {
	hobbies := []models.Hobby{
		{Name: "Run", Description: "Excercise"},
		{Name: "Study", Description: "Online courses"},
	}
	friends := []string{"Marcela", "Laura", "Constanza", "Camila"}
	person := models.Person{
		PersonID: 1,
		Names:    "Jean Carlos",
		Surnames: "Contreras Contreras",
		Age:      27,
		IsMan:    true,
		Friends:  friends,
		Hobbies:  hobbies,
	}

	utils.RenderTemplate(w, "product", person)
}
