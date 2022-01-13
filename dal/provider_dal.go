package dal

import (
	"github.com/JeanCntrs/admin-system/database"
	"github.com/JeanCntrs/admin-system/models"
)

func GetProviders() []models.Provider {
	query := "SELECT * FROM getProviders()"
	database.OpenConnection()
	rows, _ := database.Query(query)
	database.CloseConnection()

	providers := []models.Provider{}
	for rows.Next() {
		provider := models.Provider{}
		rows.Scan(&provider.ProviderId, &provider.Name, &provider.Phone, &provider.CountryName)
		providers = append(providers, provider)
	}

	return providers
}

func GetProvidersByCountryId(countryId int) []models.Provider {
	query := "SELECT * FROM getProvidersByCountryId($1)"
	database.OpenConnection()
	rows, _ := database.Query(query, countryId)
	database.CloseConnection()

	providers := []models.Provider{}
	for rows.Next() {
		provider := models.Provider{}
		rows.Scan(&provider.ProviderId, &provider.Name, &provider.Phone, &provider.CountryName)
		providers = append(providers, provider)
	}

	return providers
}

func GetProviderById(id int) models.Provider {
	query := "SELECT * FROM getProviderById($1)"
	database.OpenConnection()
	rows, _ := database.Query(query, id)
	database.CloseConnection()

	provider := models.Provider{}
	for rows.Next() {
		rows.Scan(&provider.ProviderId, &provider.Name, &provider.Address, &provider.Phone, &provider.Email, &provider.LegalRepresentative, &provider.CellPhone, &provider.CountryId, &provider.Ruc)
	}

	return provider
}
