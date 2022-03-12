package handler

import (
	"log"
	"muazwzxv/product-api/data"
	"net/http"
	"regexp"
	"strconv"
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
		regex := regexp.MustCompile(`/([0-9]+)`)
		group := regex.FindAllStringSubmatch(r.URL.Path, -1)
		if len(group) != 1 || len(group[0]) != 2 {
			http.Error(rw, "INVALID URI", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(group[0][1])
		if err != nil {
			http.Error(rw, "INVALID URI", http.StatusBadRequest)
			return
		}
		log.Printf("The id is: %d", id)

		p.updateProduct(id, rw, r)
		return
	}

	// catch
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *ProductHandler) updateProduct(id int, rw http.ResponseWriter, r *http.Request) {
	p.log.Println("Handle PUT Request")

	product := &data.Product{}
	err := product.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	// To do
	err = data.UpdateProduct(id, product)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	p.log.Printf("Updated products : %v", product)
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

func (p *ProductHandler) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.log.Println("Handle GET Request")

	products := data.GetProducts()

	err := products.ToJSON(rw)

	if err != nil {
		p.log.Fatalf("Error %s", err)
		http.Error(rw, "Unable to Marshal json", http.StatusInternalServerError)
	}
}
