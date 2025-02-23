package main

import (
	"backend/internal/api"
	"backend/pkg/database"
	"fmt"
	"log"
	"os"
)

func main() {
	// Load database configuration
	dbConfig := database.NewConfig()

	// Connect to the database
	if err := database.Connect(dbConfig); err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer func() {
		if err := database.Close(); err != nil {
			log.Printf("Error closing database connection: %v", err)
		}
	}()

	// Run migrations
	if err := database.Migrate(database.GetDB()); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Setup router
	router := api.SetupRouter()

	// Get port from environment variable
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080" // default port
	}

	// Start the server
	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
