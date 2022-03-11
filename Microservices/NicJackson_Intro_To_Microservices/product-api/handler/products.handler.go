package handler

import (
	"log"
	"muazwzxv/product-api/data"
	"net/http"
)

type ProductHandler struct {
	log *log.Logger
}

func NewProductHandler(logger *log.Logger) *ProductHandler {
	return &ProductHandler{logger}
}

func (p *ProductHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.postProduct(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		return
	}

	// catch
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *ProductHandler) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.log.Println("Handle GET Request")

	products := data.GetProducts()

	err := products.ToJSON(rw)

	if err != nil {
		p.log.Fatalf("Error %s", err)
		http.Error(rw, "Unable to Marshal json", http.StatusInternalServerError)
	}
}

func (p *ProductHandler) postProduct(rw http.ResponseWriter, r *http.Request) {
	p.log.Println("Handle POST Request")

	product := &data.Product{}

	err := product.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	data.AddProduct(product)
	p.log.Printf("Product: %v", product)
}
