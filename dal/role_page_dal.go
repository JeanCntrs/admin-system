package dal

import (
	"github.com/JeanCntrs/admin-system/database"
	"github.com/JeanCntrs/admin-system/models"
)

func GetPageByRoleTypeId(id int) []models.RolePage {
	query := "SELECT page_id FROM role_page WHERE role_type_id = $1"

	database.OpenConnection()
	rows, _ := database.Query(query, id)
	database.CloseConnection()

	rolePages := []models.RolePage{}
	for rows.Next() {
		rolePage := models.RolePage{}
		rows.Scan(&rolePage.PageId)
		rolePages = append(rolePages, rolePage)
	}

	return rolePages
}
