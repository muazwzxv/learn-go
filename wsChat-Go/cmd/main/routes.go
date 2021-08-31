package main

import (
	"github.com/bmizerany/pat"
	"net/http"
	"wsChat-Go/application/handlers"
)

func routes() http.Handler {
	route := pat.New()

	route.Get("/", http.HandlerFunc(handlers.Home))

	return route
}
