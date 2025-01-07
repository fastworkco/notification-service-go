package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// StartServer initializes and starts the HTTP server.
func StartServer(port string) error {
	// Create a new router
	router := mux.NewRouter()

	// Setup routes
	setupRoutes(router)

	// Start the server
	log.Printf("Starting server on port %s...", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Error starting server: %v", err)
		return err
	}

	return nil
}

// setupRoutes defines the API routes
func setupRoutes(router *mux.Router) {}
