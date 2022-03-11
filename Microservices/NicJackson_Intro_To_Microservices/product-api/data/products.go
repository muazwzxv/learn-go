package data

import "time"

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	SKU         string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

func GetProducts() []*Product {
	return productList
}

var productList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.50,
		SKU:         "wer453",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},

	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       2,
		SKU:         "hdj454",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
}
