package main

import (
	"log"
	"net/http"
	"wsChat-Go/application/handlers"
)

func main() {
	server := routes()
	log.Println("Serer listening on port 8080")

	log.Println("Starting channel listener")
	go handlers.ListenToSocketChannel()

	_ = http.ListenAndServe(":8080", server)

}
