package main

import (
	"log"
	"musicstore/album"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {
	router := album.NewRouter()

	//to allow access from the front-end to the methods:
	//allow CORS access. Without this, the webpage will
	//not be able to access the resources from the REST API
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})

	//Launch Server with CORS validations (https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS)
	// The handler is usually nil, which means to use DefaultServeMux.
	// Handle and HandleFunc add handlers to DefaultServeMux
	log.Fatal(http.ListenAndServe(":9000",
		handlers.CORS(allowedOrigins, allowedMethods)(router)))
}
