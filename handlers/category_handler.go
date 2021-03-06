package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/JeanCntrs/admin-system/dal"
	"github.com/JeanCntrs/admin-system/models"
	"github.com/JeanCntrs/admin-system/utils"
	"github.com/gorilla/mux"
)

type categoryForm struct {
	CategoryList []models.Category
	CategoryName string
}

func Category(w http.ResponseWriter, r *http.Request) {
	var categories []models.Category
	var searchParam string

	if r.Method == "GET" {
		categories = dal.ListCategories()
	}

	if r.Method == "POST" {
		searchParam = r.FormValue("categoryName")
		categories = dal.FilterCategories(searchParam)
	}

	category := categoryForm{CategoryList: categories, CategoryName: searchParam}

	utils.RenderTemplate(w, "category", category)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		utils.RenderTemplate(w, "create_category", nil)
	}

	if r.Method == "POST" {
		categoryId := r.FormValue("categoryId")
		categoryName := strings.TrimSpace(r.FormValue("categoryName"))
		categoryDescription := strings.TrimSpace(r.FormValue("categoryDescription"))

		category := models.Category{
			Name:        categoryName,
			Description: categoryDescription,
		}

		if categoryId == "" {
			// Insert
			errorDuplicateDataInsert := utils.ValidateDuplicateDataInsert("categories", "name", categoryName)
			if errorDuplicateDataInsert != nil {
				category.ErrorExist = true
				category.ErrorMessage = errorDuplicateDataInsert.Error()

				utils.RenderTemplate(w, "create_category", category)

				return
			}

			_, err := dal.CreateCategory(categoryName, categoryDescription)
			if err != nil {
				category.ErrorExist = true
				category.ErrorMessage = err.Error()

				utils.RenderTemplate(w, "create_category", category)

				return
			}

			http.Redirect(w, r, "/categories", http.StatusMovedPermanently)
		} else {
			// Update
			categoryIdConv, _ := strconv.Atoi(categoryId)
			category.CategoryId = categoryIdConv

			errorDuplicateDataUpdate := utils.ValidateDuplicateDataUpdate("categories", "name", categoryName, "category_id", categoryIdConv)
			if errorDuplicateDataUpdate != nil {
				category.ErrorExist = true
				category.ErrorMessage = errorDuplicateDataUpdate.Error()

				utils.RenderTemplate(w, "edit_category", category)

				return
			}

			_, err := dal.UpdateCategory(categoryIdConv, categoryName, categoryDescription)
			if err != nil {
				category.ErrorExist = true
				category.ErrorMessage = err.Error()

				utils.RenderTemplate(w, "edit_category", category)

				return
			}

			http.Redirect(w, r, "/categories", http.StatusMovedPermanently)
		}
	}
}

func EditCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	idConv, err := strconv.Atoi(id)

	if err != nil {
		panic("An error occurred")
	}

	category := dal.SearchCategoryById(idConv)

	utils.RenderTemplate(w, "edit_category", category)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	idConv, err := strconv.Atoi(id)

	if err != nil {
		panic("An error occurred")
	}

	_, errorFound := dal.DeleteCategory(idConv)

	if errorFound == nil {
		http.Redirect(w, r, "/categories", http.StatusMovedPermanently)
	}
}
