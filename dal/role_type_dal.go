package dal

import (
	"database/sql"

	"github.com/JeanCntrs/admin-system/database"
	"github.com/JeanCntrs/admin-system/models"
)

func GetRoleTypes() []models.RoleType {
	query := `SELECT role_type_id, name, description
	FROM role_type
	WHERE active = true`

	database.OpenConnection()
	rows, _ := database.Query(query)
	database.CloseConnection()

	roleTypes := []models.RoleType{}
	for rows.Next() {
		roleType := models.RoleType{}
		rows.Scan(&roleType.RoleTypeId, &roleType.Name, &roleType.Description)
		roleTypes = append(roleTypes, roleType)
	}

	return roleTypes
}

func GetRoleTypeById(id int) models.RoleType {
	query := `SELECT role_type_id, name, description
	FROM role_type
	WHERE role_type_id = $1`

	database.OpenConnection()
	rows, _ := database.Query(query, id)
	database.CloseConnection()

	roleType := models.RoleType{}
	for rows.Next() {
		rows.Scan(&roleType.RoleTypeId, &roleType.Name, &roleType.Description)
	}

	return roleType
}

func InsertRoleTypes(roleType models.RoleType, tx *sql.Tx) (int, error) {
	var roleTypeId int
	query := `INSERT INTO role_type(name, description, active) values($1, $2, true) returning role_type_id`
	err := tx.QueryRow(query, roleType.Name, roleType.Description).Scan(&roleTypeId)

	if err != nil {
		return 0, err
	}

	return roleTypeId, nil
}

func InsertRolePages(rolePage models.RolePage, tx *sql.Tx) error {
	query := `INSERT INTO role_page(role_type_id, page_id, active) values($1, $2, true)`
	_, err := tx.Exec(query, rolePage.RoleTypeId, rolePage.PageId)
	if err != nil {
		return err
	}

	return nil
}

func RegisterRoleType(roleType models.RoleType, listRolePage []models.RolePage) error {
	database.OpenConnection()

	tx, err := database.Begin()
	if err != nil {
		return err
	}

	roleTypeId, insertRoleTypeError := InsertRoleTypes(roleType, tx)

	if insertRoleTypeError != nil {
		tx.Rollback()
		return insertRoleTypeError
	}

	for _, v := range listRolePage {
		v.RoleTypeId = roleTypeId
		insertRolePageError := InsertRolePages(v, tx)
		if insertRolePageError != nil {
			tx.Rollback()
			return insertRolePageError
		}
	}

	tx.Commit()
	database.CloseConnection()
	return nil
}
