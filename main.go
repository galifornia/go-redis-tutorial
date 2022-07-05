package main

import (
	"context"
	"log"

	"github.com/galifornia/go-redis-tutorial/database"
)

func main() {
	// Open DB
	database.ConnectToRedis(context.Background())

	// Add some random key/value pairs
	database.SetToRedis(context.Background(), "bananas", "test")
	database.SetToRedis(context.Background(), "another", "test2")
	database.SetToRedis(context.Background(), "yay", "sweet")
	database.SetToRedis(context.Background(), "banana", "test3")

	// Fetch all keys that match "banana*"
	values := database.GetAllKeysFromRedis(context.Background(), "banana*")
	log.Println(values)

	// Fetch value from key pointer "yay"
	yay := database.GetFromRedis(context.Background(), "yay")
	log.Println(yay)
}
