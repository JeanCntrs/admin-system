package models

type Product struct {
	ProductId    int
	ProductName  string
	Description  string
	Price        float64
	Stock        int
	CategoryId   int
	CategoryName string
	CategoryList []Category
	Errors       []string
}
