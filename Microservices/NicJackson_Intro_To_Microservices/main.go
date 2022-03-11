package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/muazwzxv/learn-go/Microservices/NicJackson_Intro_To_Microservices/handlers"
)


func main() {

	introHandler := handlers.NewIntroHandler(log.New(os.Stdout, "intro-api", log.LstdFlags))


	serveMux := http.NewServeMux()
	serveMux.Handle("/", introHandler)

	server := &http.Server{
		Addr: ":9000",
		Handler: serveMux,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second,
	}

	server.ListenAndServe()

	/**
		http.ListenAndServe(":9000", nil)
		- Setting nil will use the default serveMux
		- Above code shows us overriding the default serveMux
		  to use our own servemux with our route handlers

		http.ListenAndServe(":9000", customServeMux)
		- customServeMux is the custom serverMux we defined
	*/
}