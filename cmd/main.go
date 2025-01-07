package main

import (
	"log"
	"notification-service-go/config"
	"notification-service-go/internal/handlers"
)

func main() {
	// Initialize config
	_ = config.LoadConfig()

	// Start HTTP server with notification handlers
	err := handlers.StartServer("8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
