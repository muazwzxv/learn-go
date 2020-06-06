package main

import (
	"fmt"
	"net/http"
	"log"
	"sync"
)

var urls = []string{
	"https://google.com",
	"https://twitter.com",
	"https://github.com",
}

var port string

func fetch(url string, wg *sync.WaitGroup) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	wg.Done()
	fmt.Println(res.Status)

	return res.Status, nil
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Homepage endpoint hit boiss")
	var wg sync.WaitGroup	

	for _, url := range urls {
		wg.Add(1)
		go fetch(url, &wg)
	}

	wg.Wait()
	fmt.Println("Returning response")
	fmt.Fprintf(w, "Response")
}

func handler(port string) {
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(port, nil))	
}

func main() {
	port := ":8080"
	handler(port)
}
