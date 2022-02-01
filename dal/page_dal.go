package dal

import (
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
