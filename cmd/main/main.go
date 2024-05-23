package main

import (
	"log"
	"net/http"

	"github/fokaz-c/go-book-manager/pkg/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)

	// Log a message indicating that the server is starting
	log.Println("Server starting on port 8080...")

	// Start the HTTP server
	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
