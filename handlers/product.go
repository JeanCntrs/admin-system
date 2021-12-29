package handlers

import (
	"net/http"
	"strconv"

	"github.com/JeanCntrs/admin-system/dal"
	"github.com/JeanCntrs/admin-system/models"
	"github.com/JeanCntrs/admin-system/utils"
	"github.com/gorilla/mux"
)

type productForm struct {
	ProductList  []models.Product
	CategoryList []models.Category
	CategoryId   string
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

	product := productForm{ProductList: productList, CategoryList: categoryList, CategoryId: searchValue}

	utils.RenderTemplate(w, "product", product)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	categoryList := dal.ListCategories()
	product := models.Product{CategoryList: categoryList}
	utils.RenderTemplate(w, "create_product", product)
}

func SaveProduct(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	idConve, _ := strconv.Atoi(id)
	name := r.FormValue("name")
	description := r.FormValue("description")
	price := r.FormValue("price")
	priceConv, _ := strconv.ParseFloat(price, 64)
	stock := r.FormValue("stock")
	stockConv, _ := strconv.Atoi(stock)
	categoryId := r.FormValue("category")
	categoryIdConv, _ := strconv.Atoi(categoryId)

	var errorMessages []string

	errorMaxName := utils.MaxLength(name, "name", 100)
	if errorMaxName != nil {
		errorMessages = append(errorMessages, errorMaxName.Error())
	}

	errorMinName := utils.MinLength(name, "name", 3)
	if errorMinName != nil {
		errorMessages = append(errorMessages, errorMinName.Error())
	}

	errorMaxDescription := utils.MaxLength(description, "description", 200)
	if errorMaxDescription != nil {
		errorMessages = append(errorMessages, errorMaxDescription.Error())
	}

	errorMinDescription := utils.MinLength(description, "description", 4)
	if errorMinDescription != nil {
		errorMessages = append(errorMessages, errorMinDescription.Error())
	}

	product := models.Product{
		ProductId:   idConve,
		ProductName: name,
		Description: description,
		Price:       priceConv,
		Stock:       stockConv,
		CategoryId:  categoryIdConv,
	}

	if id == "" {
		if len(errorMessages) > 0 {
			product.Errors = errorMessages

			utils.RenderTemplate(w, "create_product", product)

			return
		}

		categoryList := dal.ListCategories()
		product.CategoryList = categoryList

		_, err := dal.InsertProduct(product.Stock, product.CategoryId, product.ProductName, product.Description, product.Price)

		if err == nil {
			http.Redirect(w, r, "/products", http.StatusMovedPermanently)
		}
	} else {
		if len(errorMessages) > 0 {
			product.Errors = errorMessages

			utils.RenderTemplate(w, "edit_product", product)

			return
		}

		categoryList := dal.ListCategories()
		product.CategoryList = categoryList

		_, err := dal.UpdateProduct(product.ProductId, product.Stock, product.CategoryId, product.ProductName, product.Description, product.Price)

		if err == nil {
			http.Redirect(w, r, "/products", http.StatusMovedPermanently)
		}
	}
}

func EditProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	idConv, _ := strconv.Atoi(id)
	product := dal.FilterProductsById(idConv)
	categoryList := dal.ListCategories()
	product.CategoryList = categoryList
	utils.RenderTemplate(w, "edit_product", product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	idConv, err := strconv.Atoi(id)

	if err != nil {
		panic("An error occurred")
	}

	_, errorFound := dal.DeleteProduct(idConv)

	if errorFound == nil {
		http.Redirect(w, r, "/products", http.StatusMovedPermanently)
	}
}
