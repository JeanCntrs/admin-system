package handlers

import (
	"net/http"

	"github.com/JeanCntrs/admin-system/dal"
	"github.com/JeanCntrs/admin-system/models"
	"github.com/JeanCntrs/admin-system/utils"
)

func Category(w http.ResponseWriter, r *http.Request) {
	var categories = []models.Category{}

	if r.Method == "GET" {
		categories = dal.ListCategories()
	}

	if r.Method == "POST" {
		searchParam := r.FormValue("categoryName")
		categories = dal.FilterCategories(searchParam)
	}

	utils.RenderTemplate(w, "category", categories)
}
