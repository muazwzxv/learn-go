package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Article struct {
	ID          int    `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Content     string `json:"Content"`
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

func getArticle(w http.ResponseWriter, r *http.Request) {
	data := mux.Vars(r)
	id, err := strconv.Atoi(data["id"])
	if err != nil {
		errors.New("Error occured")
	}

	for _, article := range Articles {
		if article.ID == id {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func handleRequest() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home)
	router.HandleFunc("/api/Articles", getAllArticles)
	router.HandleFunc("/api/Article/{id}", getArticle)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	Articles = []Article{
		Article{1, "Awakening Of Arep", "A chinese boy waking up from sleeping", "Sikes hes gay"},
		Article{2, "Star Wars Force Awaken", "A gurl who discover a cult", "Sikes shes  gay"},
		Article{3, "Star Wars Return OF The Jedi", "A rise of an ancient cult", "Sikes hes gay"},
		Article{4, "Harry Potter Goblet Of Fire", "A nerd with magic wands", "Sikes hes gay"},
	}
	fmt.Println("Rest APi v2.0 - Mux Routers")
	handleRequest()
}
