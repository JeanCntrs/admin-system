package dal

import (
	"github.com/JeanCntrs/admin-system/database"
	"github.com/JeanCntrs/admin-system/models"
)

type CategoryList []models.Category

func ListCategories() CategoryList {
	query := `SELECT * FROM uspListCategories()`
	database.OpenConnection()
	rows, _ := database.Query(query)
	database.CloseConnection()

	categoryList := CategoryList{}
	for rows.Next() {
		category := models.Category{}
		rows.Scan(&category.Idcategoria, &category.Nombre, &category.Descripcion)
		categoryList = append(categoryList, category)
	}

	return categoryList
}
