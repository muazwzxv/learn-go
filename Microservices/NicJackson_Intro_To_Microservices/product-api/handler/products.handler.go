package handler

import (
	"log"
	"net/http"

	"github.com/muazwzxv/learn-go/Microservices/NicJackson_Intro_To_Microservices/product-api/data"
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

	if r.Method == http.MethodPut {
		return
	}

	if r.Method == http.MethodPost {
		return
	}

	// catch
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *ProductHandler) getProducts(rw http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()

	err := products.ToJSON(rw)

	if err != nil {
		p.log.Fatalf("Error %s", err)
		http.Error(rw, "Unable to Marshal json", http.StatusInternalServerError)
	}
}
