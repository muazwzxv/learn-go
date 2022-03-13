package main

import (
	"context"
	"log"
	"muazwzxv/product-api-httplib/handler"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	log := log.New(os.Stdout, "Products-api", log.LstdFlags)

	productHandler := handler.NewProductHandler(log)

	serveMux := http.NewServeMux()
	serveMux.Handle("/", productHandler)

	server := http.Server{
		Addr:         ":9000",
		Handler:      serveMux,
		ErrorLog:     log,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		log.Println("Server starting on port 9000")

		err := server.ListenAndServe()
		if err != nil {
			log.Printf("Error starting server: %s \n", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	sig := <-c
	log.Println("Received terminate request, graceful shutdown", sig)

	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	_ = server.Shutdown(timeoutContext)
}
