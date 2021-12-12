package dal

import (
	"database/sql"

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
		rows.Scan(&category.CategoryId, &category.Name, &category.Description)
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
		rows.Scan(&category.CategoryId, &category.Name, &category.Description)
		categoryList = append(categoryList, category)
	}

	return categoryList
}

func CreateCategory(name, description string) (sql.Result, error) {
	query := "INSERT INTO categoria(nombre, descripcion) values ($1, $2)"
	database.OpenConnection()
	result, err := database.Excec(query, name, description)
	database.CloseConnection()

	return result, err
}
