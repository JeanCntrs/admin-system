package dal

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/JeanCntrs/admin-system/database"
	"github.com/JeanCntrs/admin-system/models"
	"github.com/JeanCntrs/admin-system/utils"
)

func InsertUserTx(user models.User, tx *sql.Tx) error {
	encryptedPassword := utils.Encrypt(user.Password)

	query := "INSERT INTO users(username, password, person_id, role_type_id, active) values($1, $2, $3, $4, true)"
	_, err := tx.Exec(query, user.Username, encryptedPassword, user.PersonId, user.RoleTypeId)
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

func GetUserById(id int) models.User {
	query := "SELECT * FROM getUserById($1)"
	database.OpenConnection()
	rows, _ := database.Query(query, id)
	database.CloseConnection()

	user := models.User{}
	for rows.Next() {
		rows.Scan(&user.UserId, &user.Username, &user.PersonId, &user.RoleTypeId, &user.RoleTypeName, &user.Fullname)
	}

	return user
}

func UpdateUser(user models.User) (sql.Result, error) {
	query := "SELECT updateUser($1, $2, $3)"

	database.OpenConnection()
	result, err := database.Excec(query, user.UserId, user.Username, user.RoleTypeId)
	database.CloseConnection()

	return result, err
}

func DeleteUser(id int) (sql.Result, error) {
	query := "SELECT deleteUser($1)"

	database.OpenConnection()
	result, err := database.Excec(query, id)
	database.CloseConnection()

	return result, err
}

func ValidateExistingUser(username, password string) string {
	query := "SELECT COUNT(*) FROM users WHERE username = $1 AND password = $2"

	database.OpenConnection()
	quantity, err := database.QueryRow(query, username, password)
	if err != nil {
		fmt.Println("Validate User Error: ", err)
		return "0"
	}

	return strconv.Itoa(quantity)
}

func GetUserId(username, password string) string {
	query := "SELECT user_id FROM users WHERE username = $1 AND password = $2"

	database.OpenConnection()
	userId, err := database.QueryRow(query, username, password)
	if err != nil {
		fmt.Println("Validate User Error: ", err)
		return "0"
	}

	return strconv.Itoa(userId)
}
