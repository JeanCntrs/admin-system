package dal

import (
	"github.com/JeanCntrs/admin-system/database"
	"github.com/JeanCntrs/admin-system/models"
)

func ListCategories() []models.Category {
	query := `SELECT * FROM uspListCategories()`
	database.OpenConnection()
	rows, _ := database.Query(query)
	database.CloseConnection()

	categoryList := []models.Category{}
	for rows.Next() {
		category := models.Category{}
		rows.Scan(&category.Idcategoria, &category.Nombre, &category.Descripcion)
		categoryList = append(categoryList, category)
	}

	return categoryList
}

func FilterCategories(searchParam string) []models.Category {
	query := `SELECT * FROM uspFilterCategories($1)`
	database.OpenConnection()
	rows, _ := database.Query(query, searchParam)
	database.CloseConnection()

	categoryList := []models.Category{}
	for rows.Next() {
		category := models.Category{}
		rows.Scan(&category.Idcategoria, &category.Nombre, &category.Descripcion)
		categoryList = append(categoryList, category)
	}

	return categoryList
}
