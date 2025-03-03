package controller

import (
	"encoding/json"
	"goPgxSqlx/config"
	"goPgxSqlx/models"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserHandler struct {
  DB *sqlx.DB
}

// GetAllUsers returns all users
func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
  users := []models.User{}
  err := h.DB.SelectContext(r.Context(), &users, "SELECT * FROM users")
  if err != nil {
    log.Printf("Error getting users: %v", err)
    http.Error(w, config.JsonError("Failed to retrieve users"), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(users)
}

// GetUserByID returns a specific user by ID
func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
  id := chi.URLParam(r, "id")
  userID, err := uuid.Parse(id)
  if err != nil {
    http.Error(w, config.JsonError("Invalid user ID"), http.StatusBadRequest)
    return
  }

  user := models.User{}
  err = h.DB.GetContext(r.Context(), &user, "SELECT * FROM users WHERE id = $1", userID)
  if err != nil {
    log.Printf("Error getting user %s: %v", userID, err)
    http.Error(w, config.JsonError("User not found"), http.StatusNotFound)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(user)
}
