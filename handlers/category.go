package handlers

import (
	"net/http"

	"github.com/JeanCntrs/admin-system/dal"
	"github.com/JeanCntrs/admin-system/utils"
)

func Category(w http.ResponseWriter, r *http.Request) {
	categoryList := dal.ListCategories()

	utils.RenderTemplate(w, "category", categoryList)
}
