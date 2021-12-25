package models

type Product struct {
	ProductId    int
	ProductName  string
	Price        float64
	Stock        int
	CategoryName string
	CategoryList []Category
}
