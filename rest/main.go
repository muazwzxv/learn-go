package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)


type Article struct {
	Title string `json:"Title"`
	Description string `json:"Description"`
	Content string `json:"content"`
}


// Create a global Article array
// Store data in memory, emulate effect of a database
var Articles []Article


func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Hompage boi")
	fmt.Println("Endpoint: Return the homepage")
}

func getAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint: Return All Articles")
	json.NewEncoder(w).Encode(Articles)
}

func handleRequest() {
	http.HandleFunc("/", home)
	http.HandleFunc("/api/Articles", getAllArticles)
	log.Fatal(http.ListenAndServe(":8080", nil))
}


func main() {
	Articles = []Article {
		Article{"Awakening Of Arep", "A chinese boy waking up from sleeping", "Sikes hes gay"},
		Article{"Star Wars Force Awaken", "A gurl who discover a cult", "Sikes shes  gay"},
		Article{"Star Wars Return OF The Jedi", "A rise of an ancient cult", "Sikes hes gay"},
		Article{"Harry Potter Goblet Of Fire", "A nerd with magic wands", "Sikes hes gay"},
	}
	handleRequest()
}
