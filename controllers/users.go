package controllers

import (
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
    w.WriteHeader(http.StatusInternalServerError)
    config.WriteJSONResponse(w, "Failed to retrieve users", http.StatusInternalServerError)
    return
  }

  config.WriteJSONResponse(w, users, http.StatusOK)
}

// GetUserByID returns a specific user by ID
func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
  id := chi.URLParam(r, "id")
  userID, err := uuid.Parse(id)
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    config.WriteJSONResponse(w, "Invalid user ID", http.StatusBadRequest)
    return
  }

  user := models.User{}
  err = h.DB.GetContext(r.Context(), &user, "SELECT * FROM users WHERE id = $1", userID)
  if err != nil {
    log.Printf("Error getting user %s: %v", userID, err)
    config.WriteJSONResponse(w, "User not found", http.StatusNotFound)
    return
  }

  config.WriteJSONResponse(w, user, http.StatusOK)
}
