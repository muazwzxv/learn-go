package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/muazwzxv/learn-go/Microservices/NicJackson_Intro_To_Microservices/handlers"
)

func main() {

	logger := log.New(os.Stdout, "intro-api", log.LstdFlags)
	introHandler := handlers.NewIntroHandler(logger)

	serveMux := http.NewServeMux()
	serveMux.Handle("/", introHandler)

	server := &http.Server{
		Addr:         ":9000",
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, syscall.SIGTERM)

	sig := <-signalChannel
	logger.Println("Received terminate request, graceful shutdown", sig)

	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(timeoutContext)
}

/**
http.ListenAndServe(":9000", nil)
- Setting nil will use the default serveMux
- Above code shows us overriding the default serveMux
  to use our own servemux with our route handlers

http.ListenAndServe(":9000", customServeMux)
- customServeMux is the custom serverMux we defined
*/
