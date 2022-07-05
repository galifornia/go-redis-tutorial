package database

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var DB *redis.Client

// ConnectToRedis: opens Redis database connection
func ConnectToRedis(ctx context.Context) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(pong)

	DB = client
}

// SetToRedis: sets key/value pair in DB
func SetToRedis(ctx context.Context, key, val string) {
	err := DB.Set(ctx, key, val, 0).Err()
	if err != nil {
		log.Println(err)
	}
}

// GetFromRedis: gets value from key
func GetFromRedis(ctx context.Context, key string) string {
	val, err := DB.Get(ctx, key).Result()
	if err != nil {
		log.Println(err)
	}
	return val
}

// GetAllKeysFromRedis: get keys that match regular expression
func GetAllKeysFromRedis(ctx context.Context, regex string) []string {
	keys := []string{}

	iter := DB.Scan(ctx, 0, regex, 0).Iterator()
	for iter.Next(ctx) {
		keys = append(keys, iter.Val())
	}

	if err := iter.Err(); err != nil {
		log.Println(err)
	}

	return keys
}
