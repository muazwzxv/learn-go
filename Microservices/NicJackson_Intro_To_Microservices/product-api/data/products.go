package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	SKU         string    `json:"sku"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
	DeletedAt   time.Time `json:"-"`
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	// This writes the encoded json to w
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func GetProducts() Products {
	return productList
}

// Dummy data
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
