package handlers

import (
	"net/http"
	"strconv"

	"github.com/JeanCntrs/admin-system/dal"
	"github.com/JeanCntrs/admin-system/models"
	"github.com/JeanCntrs/admin-system/utils"
)

type productForm struct {
	ProductList  []models.Product
	CategoryList []models.Category
}

func Product(w http.ResponseWriter, r *http.Request) {
	var productList []models.Product
	categoryList := dal.ListCategories()

	searchValue := r.FormValue("categoryId")

	if r.Method == "GET" {
		productList = dal.ListProducts()
	}

	if r.Method == "POST" {
		if searchValue == "" {
			productList = dal.ListProducts()
		} else {
			categoryId, err := strconv.Atoi(searchValue)
			if err != nil {
				panic("Category id cannot be converted")
			}

			productList = dal.FilterProductsByCategory(categoryId)
		}
	}

	product := productForm{ProductList: productList, CategoryList: categoryList}

	utils.RenderTemplate(w, "product", product)
}
