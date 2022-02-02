package dal

import (
	"database/sql"

	"github.com/JeanCntrs/admin-system/database"
	"github.com/JeanCntrs/admin-system/models"
)

func GetPages() []models.Page {
	query := "SELECT * FROM getPages()"
	database.OpenConnection()
	rows, _ := database.Query(query)
	database.CloseConnection()

	pages := []models.Page{}
	for rows.Next() {
		page := models.Page{}
		rows.Scan(&page.PageId, &page.Message, &page.Route)
		pages = append(pages, page)
	}

	return pages
}

func GetPageById(id int) models.Page {
	query := "SELECT * FROM getPageById($1)"
	database.OpenConnection()
	rows, _ := database.Query(query, id)
	database.CloseConnection()

	page := models.Page{}
	for rows.Next() {
		rows.Scan(&page.PageId, &page.Message, &page.Route)
	}

	return page
}

func InsertPage(page models.Page) (sql.Result, error) {
	query := "SELECT insertPage($1, $2)"

	database.OpenConnection()
	result, err := database.Excec(query, page.Message, page.Route)
	database.CloseConnection()

	return result, err
}

func UpdatePage(page models.Page) (sql.Result, error) {
	query := "SELECT updatePage($1, $2, $3)"

	database.OpenConnection()
	result, err := database.Excec(query, page.PageId, page.Message, page.Route)
	database.CloseConnection()

	return result, err
}
