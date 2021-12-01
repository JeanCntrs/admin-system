package handlers

import (
	"net/http"

	"github.com/JeanCntrs/admin-system/dal"
	"github.com/JeanCntrs/admin-system/utils"
)

func Product(w http.ResponseWriter, r *http.Request) {
	productList := dal.ListProducts()

	utils.RenderTemplate(w, "product", productList)
}
