package main

import (
	"fmt"
	"net/http"
	"log"
)

var urls = []string{
	"https://google.com",
	"https://twitter.com",
	"https://github.com",
}

func fetch(url string) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Status)
}

func home(w http.ResponseWriter, r *http.Response) {
	fmt.Println("Homepage endpoint hit boiss")

	for _, url := range urls {
		go fetch(url)
	}
	fmt.Println("Returning response")
	fmt.Fprintf(w, "All response received")
}

func handler() {
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(port, nil))
}

func init() {
	port := ":8080"
}

func main() {
	handler()
}
