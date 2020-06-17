package main

import (
	"fmt"
	"log"
	"net/http"
)


type Article struct {
	Title string `json:"Title"`
	Description string `json:"Desc"`
	Content string `json:"content"`
}


// Create a global Article array
// Store data in memory, emulate effect of a database
var Articles []Article


func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Hompage boi")
	fmt.Println("Main Endpoint")
}

func handleRequest() {
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	Articles = []Article {
		Article{"Awakening Of Arep", "A chinese boy waking up from sleeping", "Sikes hes gay"},
	}
	handleRequest()
}
