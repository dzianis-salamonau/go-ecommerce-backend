package products

import (
	"encoding/json"
	"net/http"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	CategoryID  int     `json:"category_id"`
}

// Dummy data
var products = []Product{
	{ID: 1, Name: "Laptop", Description: "A high-end gaming laptop", Price: 1500.50, Stock: 10, CategoryID: 1},
	{ID: 2, Name: "Mouse", Description: "Wireless mouse", Price: 25.00, Stock: 200, CategoryID: 2},
}

// Handler serves HTTP requests for the /products endpoint
func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
