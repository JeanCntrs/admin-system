package handlers

import (
	"net/http"

	"github.com/JeanCntrs/admin-system/dal"
	"github.com/JeanCntrs/admin-system/utils"
)

func Category(w http.ResponseWriter, r *http.Request) {
	person := dal.ListCategories()

	utils.RenderTemplate(w, "category", person)
}
