package dal

import (
	"github.com/JeanCntrs/admin-system/database"
	"github.com/JeanCntrs/admin-system/models"
)

func GetCountries() []models.Country {
	query := `SELECT * FROM getCountries()`
	database.OpenConnection()
	rows, _ := database.Query(query)
	database.CloseConnection()

	countries := []models.Country{}
	for rows.Next() {
		country := models.Country{}
		rows.Scan(&country.CountryId, &country.Name, &country.Capital)
		countries = append(countries, country)
	}

	return countries
}
