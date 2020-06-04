package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var port string
var counter int
var mutex = &sync.Mutex{}

func init() {
	port = ":8080"
}

func printGreet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from the world of golang server, %q", html.EscapeString(r.URL.Path))
}

func increment(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprintf(w, "incrementing ..... ", strconv.Itoa(counter), html.EscapeString(r.URL.Path))
	mutex.Unlock()
}

func main() {

	http.HandleFunc("/", printGreet)
	http.HandleFunc("/increment", increment)

	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Greet madafaka, %q", html.EscapeString(r.URL.Path))
	})

	log.Println("Server is listening on", port)
	log.Fatal(http.ListenAndServe(port, nil))
	log.Fatal(http.ListenAndServe(port, nil))

}

// fmt.Fprintf(io writer, string format, interface) (int, err)
// Returns number of byte written and error if exist

/* Mutex locking
- Mutex locking will lock a specific resources and let one process use it a time, pending request 
will be queue
*/

