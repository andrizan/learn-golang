package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

// Cache is a global Redis client instance
var Cache *redis.Client

// InitRedis initializes the Redis connection with configuration
func InitRedis() {
		redisUrl := os.Getenv("REDIS_URL")
    if redisUrl == "" {
        log.Fatal("REDIS_URL is not set")
    }
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Create Redis client
    rdb := redis.NewClient(&redis.Options{
        Addr:         redisUrl,
        Password:     os.Getenv("REDIS_PASSWORD"),
        DB:          0,
        PoolSize:     10,
        MinIdleConns: 5,
        MaxRetries:   3,
        DialTimeout:  5 * time.Second,
        ReadTimeout:  3 * time.Second,
        WriteTimeout: 3 * time.Second,
        PoolTimeout:  4 * time.Second,
    })

    // Test connection
    if err := rdb.Ping(ctx).Err(); err != nil {
        log.Fatalf("Could not connect to Redis: %v", err)
    }

    Cache = rdb
    fmt.Println("Redis connected successfully")
}

// CloseRedis closes the Redis connection
func CloseRedis() {
    if Cache != nil {
        if err := Cache.Close(); err != nil {
            log.Printf("Error closing Redis connection: %v", err)
        }
    }
}
