package main

import (
	"log"
	"net/http"

	"github.com/dzianis-salamonau/go-ecommerce-backend/internal/db"
	"github.com/dzianis-salamonau/go-ecommerce-backend/pkg/users"
)

func main() {
	// Initialize database connection
	database, err := db.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer database.Close()

	// Set the DB connection in the users package
	users.DB = database

	// Initialize routes
	http.HandleFunc("/users", users.Handler)

	// Start server
	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
