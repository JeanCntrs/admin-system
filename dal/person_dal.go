package dal

import (
	"database/sql"

	"github.com/JeanCntrs/admin-system/database"
	"github.com/JeanCntrs/admin-system/models"
)

func GetPersons() []models.Person {
	query := "SELECT * FROM getPersons()"
	database.OpenConnection()
	rows, _ := database.Query(query)
	database.CloseConnection()

	persons := []models.Person{}
	for rows.Next() {
		person := models.Person{}
		rows.Scan(&person.PersonId, &person.Fullname, &person.NameTypePerson, &person.Birthday)
		person.FormattedBirthday = person.Birthday.Format("02/01/2006")
		persons = append(persons, person)
	}

	return persons
}

func GetPersonsByFullname(fullname string) []models.Person {
	query := "SELECT * FROM getPersonsByFullname($1)"
	database.OpenConnection()
	rows, _ := database.Query(query, fullname)
	database.CloseConnection()

	persons := []models.Person{}
	for rows.Next() {
		person := models.Person{}
		rows.Scan(&person.PersonId, &person.Fullname, &person.NameTypePerson, &person.Birthday)
		person.FormattedBirthday = person.Birthday.Format("02/01/2006")
		persons = append(persons, person)
	}

	return persons
}

func GetPersonById(id int) models.Person {
	query := "SELECT * FROM getPersonById($1)"
	database.OpenConnection()
	rows, _ := database.Query(query, id)
	database.CloseConnection()

	person := models.Person{}
	for rows.Next() {
		rows.Scan(&person.PersonId, &person.Name, &person.FatherLastName, &person.MotherLastName, &person.TypePersonId, &person.Birthday)
		person.FormattedBirthday = person.Birthday.Format("02/01/2006")
	}

	return person
}

func GetTypePersons() []models.TypePerson {
	query := "SELECT * FROM getTypePersons()"
	database.OpenConnection()
	rows, _ := database.Query(query)
	database.CloseConnection()

	typePersons := []models.TypePerson{}
	for rows.Next() {
		typePerson := models.TypePerson{}
		rows.Scan(&typePerson.TypePersonId, &typePerson.Name)
		typePersons = append(typePersons, typePerson)
	}

	return typePersons
}

func InsertPerson(person models.Person) (sql.Result, error) {
	query := "SELECT insertPerson($1, $2, $3, $4, $5)"

	database.OpenConnection()
	result, err := database.Excec(query, person.Name, person.FatherLastName, person.MotherLastName, person.TypePersonId, person.Birthday)
	database.CloseConnection()

	return result, err
}

func UpdatePerson(person models.Person) (sql.Result, error) {
	query := "SELECT updatePerson($1, $2, $3, $4, $5, $6)"

	database.OpenConnection()
	result, err := database.Excec(query, person.PersonId, person.Name, person.FatherLastName, person.MotherLastName, person.TypePersonId, person.Birthday)
	database.CloseConnection()

	return result, err
}
