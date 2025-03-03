package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	config "goPgxSqlx/config"
	"goPgxSqlx/controller"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/joho/godotenv/autoload" // Load .env file
)

func main() {

  // Inisialisasi database
  db := config.DbHandler{}
  db.InitDB()
  defer db.CloseDB()

  config.InitRedis()
  defer config.CloseRedis()

  // Create Chi router
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

  // Create Handler instance
  userHandler := &controller.UserHandler{DB: db.DB}

	// Routes
	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Get("/", userHandler.GetAllUsers)
			r.Get("/{id}", userHandler.GetUserByID)
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
