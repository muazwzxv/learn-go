package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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
	err := formatJSON()
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(Articles)
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint: Return Article By ID")
	data := mux.Vars(r)

	// Convert string to ints
	// Url is already string, no need for error checking
	id, _ := strconv.Atoi(data["id"])

	// Fetch article from datastore
	// Anonymous function implementation
	article, err := func(id int) (*Article, error) {
		for _, article := range Articles {
			if article.ID == id {
				return &article, nil
			}
		}
		return nil, errors.New("Not found")

	}(id)
	//article, err := getArticleHanlder(id)
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(article)
}

/*func getArticleHanlder(id int) (*Article, error) {
	for _, article := range Articles {
		if article.ID == id {
			return &article, nil
		}
	}
	return nil, errors.New("Not found")
}*/

func createArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint: Post new article")
	// Fetch the req body from the request n unmarshall into the Article struct
	body, _ := ioutil.ReadAll(r.Body)

	var article Article
	json.Unmarshal(body, &article)

	// Update the global Article array to include the new data
	Articles = append(Articles, article)

	err := formatJSON()
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(article)
}

func deleteArticleByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint: Delete article by id")

	data := mux.Vars(r)

	id, _ := strconv.Atoi(data["id"])

	for i, article := range Articles {
		if article.ID == id {
			Articles = append(Articles[:i], Articles[i+1:]...)
		}
	}

	err := formatJSON()
	if err != nil {
		fmt.Println(err)
	}
}

func handleRequest() {

	//router := mux.NewRouter().StrictSlash(true)
	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/api/articles", getAllArticles)
	router.HandleFunc("/api/article/{id}", getArticle).Methods("GET")
	router.HandleFunc("/api/article", createArticle).Methods("POST")
	router.HandleFunc("/api/article/{id}", deleteArticleByID).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":3000", router))
}

func formatJSON() error {
	byteArray, err := json.MarshalIndent(Articles, "", " ")
	if err != nil {
		return errors.New("Failed to convert to json")
	}
	fmt.Println(string(byteArray))
	return nil
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
