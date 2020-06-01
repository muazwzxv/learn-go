package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

var port string

func init() {
	port = ":8080"
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from the world of golang server, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Greet madafaka, %q", html.EscapeString(r.URL.Path))
	})

	log.Println("Server is listening on", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

// fmt.Fprintf(io writer, string format, interface) (int, err)
// Returns number of byte written and error if exist
