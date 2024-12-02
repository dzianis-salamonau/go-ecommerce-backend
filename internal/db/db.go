package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// Connect establishes a connection to the database using the DATABASE_URL environment variable.
func Connect() (*sql.DB, error) {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the database URL from environment variable
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return nil, fmt.Errorf("DATABASE_URL environment variable not set")
	}

	// Open a connection to the database
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, fmt.Errorf("could not open database: %v", err)
	}

	// Ping the database to ensure the connection is valid
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("could not connect to database: %v", err)
	}

	// Return the valid connection
	return db, nil
}
