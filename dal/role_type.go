package dal

import (
	"github.com/JeanCntrs/admin-system/database"
	"github.com/JeanCntrs/admin-system/models"
)

func GetRoleTypes() []models.RoleType {
	query := `SELECT role_type_id, name
	FROM role_type
	WHERE active = true`

	database.OpenConnection()
	rows, _ := database.Query(query)
	database.CloseConnection()

	roleTypes := []models.RoleType{}
	for rows.Next() {
		roleType := models.RoleType{}
		rows.Scan(&roleType.RoleTypeId, &roleType.Name)
		roleTypes = append(roleTypes, roleType)
	}

	return roleTypes
}
