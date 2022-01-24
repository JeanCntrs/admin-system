package dal

import (
	"database/sql"

	"github.com/JeanCntrs/admin-system/models"
)

func InsertUser(user models.User, tx *sql.Tx) error {
	query := "INSERT INTO users(username, password, person_id, role_type_id, active) values($1, $2, $3, $4, true)"
	_, err := tx.Exec(query, user.Username, user.Password, user.PersonId, user.RoleTypeId)
	if err != nil {
		return err
	}

	return nil
}

func UpdateUser(personId int, tx *sql.Tx) error {
	query := "UPDATE persons SET registered_user = true WHERE person_id = $1"
	_, err := tx.Exec(query, personId)
	if err != nil {
		return err
	}

	return nil
}
