package handler

import (
	"encoding/json"
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

func (p *ProductHandler) ServeHTTP(wr http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()

	if data, err := json.Marshal(products); err != nil {
		p.log.Fatalf("Error %s", err)
		http.Error(wr, "Unable to Marshal json", http.StatusInternalServerError)
	} else {
		wr.WriteHeader(http.StatusAccepted)
		wr.Write(data)
	}

}
