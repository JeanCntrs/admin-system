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

func GetUsers() []models.User {
	query := `SELECT u.user_id, u.username, p.name||' '||p.father_last_name||' '||p.mother_last_name fullname, rt.name
	FROM users u INNER JOIN persons p
	ON p.person_id = u.person_id
	INNER JOIN role_type rt
	ON rt.role_type_id = u.role_type_id
	WHERE u.active = true`

	database.OpenConnection()
	rows, _ := database.Query(query)
	database.CloseConnection()

	users := []models.User{}
	for rows.Next() {
		user := models.User{}
		rows.Scan(&user.UserId, &user.Username, &user.Fullname, &user.RoleTypeName)
		users = append(users, user)
	}

	return users
}
