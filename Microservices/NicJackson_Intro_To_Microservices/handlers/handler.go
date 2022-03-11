package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type IntroHandler struct {
	logger *log.Logger
}

func NewIntroHandler(logger *log.Logger) *IntroHandler {
	return &IntroHandler{logger}
}

func (h *IntroHandler) ServeHTTP(wr http.ResponseWriter, r http.Request) {
	log.Print("Hello Mate")

	if data, err := ioutil.ReadAll(r.Body); err != nil {
		http.Error(wr, "Error", http.StatusBadRequest)
	} else {
		fmt.Fprintf(wr, "Hello %s", data)
	}
}