package main

import(
	"fmt"
	"log"
	"net/http"
	"html"
)

var port string

func init() {
	port = ":8080"
}
func main() {
	
	// http.Handle is similar to http.HandleFunc, they are handler for routes
	
	// http.Handle takes the string path and a handler 
	
	// http.HandleFunc takes the path and a function with (http.ResponseWriter and *http.Request) as
	// the argument
	
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/hehe", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "The url is,  %q", html.EscapeString(r.URL.Path[1:]))
	})
	
	fmt.Println("Server is running on port ",port)
	log.Fatal(http.ListenAndServe(port, nil))
}
