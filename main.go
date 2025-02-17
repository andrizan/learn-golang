package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	config "goPgxSqlx/config"
	"goPgxSqlx/models"

	_ "github.com/joho/godotenv/autoload" // Load .env file
)

func main() {

	// Inisialisasi database
	config.InitDB()
	defer config.CloseDB()

	config.InitRedis()
	defer config.CloseRedis()

	// Contoh query
	var count int
	err := config.DB.Get(&count, "SELECT COUNT(id) FROM users")
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}
	fmt.Printf("Total users: %d\n", count)

	var users []models.User
	err = config.DB.Select(&users, "SELECT id, name, email FROM users")
	if err != nil {
		log.Fatalf("Error fetching users: %v", err)
	}

	for _, user := range users {
		userJSON, err := json.Marshal(user)
		if err != nil {
			log.Fatalf("Error marshalling user to JSON: %v", err)
		}
		fmt.Println(string(userJSON))
	}

	ctx := context.Background()
	val, err := config.Cache.Get(ctx, "data").Result()
	if err != nil {
		log.Fatalf("redis error: %v", err)
	}
	fmt.Println(val)
}
