package service

import (
	"fmt"
	"testing"
)

func TestRedis(t *testing.T) {
	InitClient()
	client := redisClient
	// Ping the Redis server to check the connection
	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println("Failed to connect to Redis:", err)
		return
	}
	fmt.Println("Connected to Redis:", pong)

	// Set a key-value pair
	err = client.Set("exampleKey", "exampleValue", 0).Err()
	if err != nil {
		fmt.Println("Failed to set key:", err)
		return
	}

	// Close the Redis client when done
	err = client.Close()
	if err != nil {
		fmt.Println("Failed to close Redis client:", err)
		return
	}
}
