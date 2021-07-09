package main

import (
	"github.com/go-redis/redis"
	"key-value-storage/storage"
)

func main() {
	redisClient := storage.Client()
	defer func(redisConn *redis.Client) {
		storage.CheckErr(redisConn.Close())
	}(redisClient)
}
