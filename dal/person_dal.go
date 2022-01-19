package dal

import (
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
		persons = append(persons, person)
	}

	return persons
}
