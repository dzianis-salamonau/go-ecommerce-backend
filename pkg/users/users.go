package users

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

// User represents a user in the system
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Password string `json:"-"`
}

// DB is a global database connection
var DB *sql.DB

// FetchUsers retrieves all users from the database
func FetchUsers() ([]User, error) {
	rows, err := DB.Query("SELECT id, name, email, role FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Role); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// Handler serves HTTP requests for the /users endpoint
func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	users, err := FetchUsers()
	if err != nil {
		log.Printf("Error fetching users: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		log.Printf("Error encoding users to JSON: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
