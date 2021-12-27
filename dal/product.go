package dal

import (
	"database/sql"

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
		rows.Scan(&product.ProductId, &product.ProductName, &product.Price, &product.Stock, &product.CategoryName)
		productList = append(productList, product)
	}

	return productList
}

func FilterProductsByCategory(categoryId int) []models.Product {
	query := "SELECT * FROM uspFilterProductsByCategory($1)"
	database.OpenConnection()
	rows, _ := database.Query(query, categoryId)
	database.CloseConnection()

	productList := []models.Product{}
	for rows.Next() {
		product := models.Product{}
		rows.Scan(&product.ProductId, &product.ProductName, &product.Price, &product.Stock, &product.CategoryName)
		productList = append(productList, product)
	}

	return productList
}

func DeleteProduct(id int) (sql.Result, error) {
	query := "SELECT uspDeleteProduct($1)"

	database.OpenConnection()
	result, err := database.Excec(query, id)
	database.CloseConnection()

	return result, err
}

func UpdateProduct(id, stock, categoryId int, name, description string, price float64) (sql.Result, error) {
	query := "SELECT uspUpdateProduct($1, $2, $3, $4, $5, $6)"

	database.OpenConnection()
	result, err := database.Excec(query, id, name, description, price, stock, categoryId)
	database.CloseConnection()

	return result, err
}

func InsertProduct(stock, categoryId int, name, description string, price float64) (sql.Result, error) {
	query := "SELECT uspInsertProduct($1, $2, $3, $4, $5)"

	database.OpenConnection()
	result, err := database.Excec(query, name, description, price, stock, categoryId)
	database.CloseConnection()

	return result, err
}

func FilterProductsById(id int) models.Product {
	query := "SELECT * FROM uspGetProductById($1)"
	database.OpenConnection()
	rows, _ := database.Query(query, id)
	database.CloseConnection()

	product := models.Product{}
	for rows.Next() {
		rows.Scan(&product.ProductId, &product.ProductName, &product.Description, &product.Price, &product.Stock, &product.CategoryId)
	}

	return product
}
