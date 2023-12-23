package service

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

//	type Storage struct {
//		redisClient *redis.Client
//	}
var redisClient *redis.Client

const DURATION = 6 * time.Hour

func InitClient() {
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379", // Replace with your Redis server address
	})
	redisClient = client
}

func SetUrl(originUrl string, shortUrl string) error {
	// Set a key-value pair
	err := redisClient.Set(shortUrl, originUrl, DURATION).Err()
	if err != nil {
		fmt.Println("Failed to set key:", err)
	}
	return err
}

func GetUrl(shortUrl string) (value string, err error) {
	// Get the value for a key
	value, err = redisClient.Get(shortUrl).Result()
	if err != nil {
		fmt.Println("Failed to get key:", err)
		return
	}
	return
}
