package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	config "goPgxSqlx/config"
	"goPgxSqlx/models"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	_ "github.com/joho/godotenv/autoload" // Load .env file
)

func main() {

  // Inisialisasi database
  config.InitDB()
  defer config.CloseDB()

  config.InitRedis()
  defer config.CloseRedis()

  // Create Chi router
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

	// Routes
	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Get("/", GetAllUsers)
			r.Get("/{id}", GetUserByID)
		})
	})

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	fmt.Printf("Server starting on port %s...\n", port)
	if err := http.ListenAndServe("localhost:"+port, r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

// GetAllUsers returns all users
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := []models.User{}
	err := config.DB.SelectContext(r.Context(), &users, "SELECT * FROM users")
	if err != nil {
		log.Printf("Error getting users: %v", err)
		http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// GetUserByID returns a specific user by ID
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	userID, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user := models.User{}
	err = config.DB.GetContext(r.Context(), &user, "SELECT * FROM users WHERE id = $1", userID)
	if err != nil {
		log.Printf("Error getting user %s: %v", userID, err)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
