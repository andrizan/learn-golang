package models

import "time"

type User struct {
	ID           string    `db:"id" json:"id"`
	Name         string    `db:"name" json:"name"`
	Email        string    `db:"email" json:"email"`
	PasswordHash string    `db:"password_hash" json:"-"`
	UserRole     string    `db:"user_role" json:"user_role"`
	UserStatus   int       `db:"user_status" json:"user_status"`
	UpdatedAt    *time.Time `db:"updated_at" json:"updated_at"`
	CreatedAt    *time.Time `db:"created_at" json:"created_at"`
}
