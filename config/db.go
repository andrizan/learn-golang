package config

import (
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib" // Register pgx driver
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitDB() {
    dsn := os.Getenv("DATABASE_URL")
    if dsn == "" {
        log.Fatal("DATABASE_URL is not set")
    }

    // Create database connection with pgx driver
    db, err := sqlx.Connect("pgx", dsn)
    if err != nil {
        log.Fatalf("Unable to connect to database: %v", err)
    }

    // Configure connection pool
    db.SetMaxOpenConns(25)
    db.SetMaxIdleConns(5)
    db.SetConnMaxLifetime(30 * time.Minute)
    db.SetConnMaxIdleTime(10 * time.Minute)

    // Test connection
    if err := db.Ping(); err != nil {
        log.Fatalf("Cannot ping database: %v", err)
    }

    DB = db
    fmt.Println("Database connected successfully with connection pool")
}

// CloseDB closes the database connection
func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
