package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/malfazakki/go-blog/config"
	"github.com/malfazakki/go-blog/routes"
)

func main() {
	// Initialize database connection
	config.ConnectDatabase()

	// Setup routes
	router := routes.SetupRoutes(config.DB)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}

	log.Printf("Server running on port %s...", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
