package main

import (
	"log"
	"net/http"
)

func main() {
	server := routes()

	log.Println("Serer listening on port 8080")

	_ = http.ListenAndServe(":8080", server)

}
