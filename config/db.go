package config

import (
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib" // Register pgx driver
	"github.com/jmoiron/sqlx"
)

type DbHandler struct {
  DB *sqlx.DB
}

func (db *DbHandler) InitDB() {
  dsn := os.Getenv("DATABASE_URL")
  if dsn == "" {
    log.Fatal("DATABASE_URL is not set")
  }

  // Create database connection with pgx driver
  conn, err := sqlx.Connect("pgx", dsn)
  if err != nil {
    log.Fatalf("Unable to connect to database: %v", err)
  }

  // Configure connection pool
  conn.SetMaxOpenConns(25)
  conn.SetMaxIdleConns(5)
  conn.SetConnMaxLifetime(30 * time.Minute)
  conn.SetConnMaxIdleTime(10 * time.Minute)

  // Test connection
  if err := conn.Ping(); err != nil {
    log.Fatalf("Cannot ping database: %v", err)
  }

  db.DB = conn
  fmt.Println("Database connected successfully with connection pool")
}

// CloseDB closes the database connection
func (db *DbHandler) CloseDB() {
  if db.DB != nil {
    db.DB.Close()
  }
}
