package orders

import (
	"encoding/json"
	"net/http"
)

type Order struct {
	ID         int     `json:"id"`
	UserID     int     `json:"user_id"`
	TotalPrice float64 `json:"total_price"`
	Status     string  `json:"status"`
}

// Dummy data
var orders = []Order{
	{ID: 1, UserID: 1, TotalPrice: 1525.50, Status: "completed"},
	{ID: 2, UserID: 2, TotalPrice: 50.00, Status: "pending"},
}

// Handler serves HTTP requests for the /orders endpoint
func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}
