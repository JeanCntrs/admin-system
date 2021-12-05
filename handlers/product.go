package handlers

import (
	"net/http"

	"github.com/JeanCntrs/admin-system/dal"
	"github.com/JeanCntrs/admin-system/models"
	"github.com/JeanCntrs/admin-system/utils"
)

type productForm struct {
	ProductList  []models.Product
	CategoryList []models.Category
}

func Product(w http.ResponseWriter, r *http.Request) {
	productList := dal.ListProducts()
	categoryList := dal.ListCategories()

	product := productForm{ProductList: productList, CategoryList: categoryList}

	utils.RenderTemplate(w, "product", product)
}
