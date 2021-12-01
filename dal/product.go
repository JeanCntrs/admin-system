package dal

import (
	"github.com/JeanCntrs/admin-system/database"
	"github.com/JeanCntrs/admin-system/models"
)

func ListProducts() []models.Product {
	query := "SELECT * FROM uspListProducts()"
	database.OpenConnection()
	rows, _ := database.Query(query)
	database.CloseConnection()

	productList := []models.Product{}
	for rows.Next() {
		product := models.Product{}
		rows.Scan(&product.IdProduct, &product.ProductName, &product.Price, &product.Stock, &product.CategoryName)
		productList = append(productList, product)
	}

	return productList
}
