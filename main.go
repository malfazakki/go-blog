package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/malfazakki/go-blog/config"
	"github.com/malfazakki/go-blog/routes"
	"github.com/malfazakki/go-blog/seeders"
)

func main() {
	// Parse command line flags
	seed := flag.Bool("seed", false, "Seed the database with sample data")
	flag.Parse()

	// Initialize database connection
	config.ConnectDatabase()

	// Run seeders it the seed flag is provided
	if *seed {
		log.Println("Seeding database...")
		seeders.RunSeeders(config.DB)
		log.Println("Database seeding completed")
		return
	}

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
