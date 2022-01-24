package dal

import (
	"database/sql"

	"github.com/JeanCntrs/admin-system/database"
	"github.com/JeanCntrs/admin-system/models"
)

func InsertUserTx(user models.User, tx *sql.Tx) error {
	query := "INSERT INTO users(username, password, person_id, role_type_id, active) values($1, $2, $3, $4, true)"
	_, err := tx.Exec(query, user.Username, user.Password, user.PersonId, user.RoleTypeId)
	if err != nil {
		return err
	}

	return nil
}

func UpdateUserTx(personId int, tx *sql.Tx) error {
	query := "UPDATE persons SET registered_user = true WHERE person_id = $1"
	_, err := tx.Exec(query, personId)
	if err != nil {
		return err
	}

	return nil
}

func RegisterUserTx(user models.User) error {
	database.OpenConnection()
	tx, err := database.Begin()
	if err != nil {
		return err
	}

	err = UpdateUserTx(user.PersonId, tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = InsertUserTx(user, tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	database.CloseConnection()

	return nil
}
