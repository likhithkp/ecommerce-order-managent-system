package redis

import (
	"context"
	"log"
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
	client *redis.Client
	once   sync.Once
)

func Redis() *redis.Client {
	once.Do(func() {
		client = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})

		if err := client.Ping(context.Background()).Err(); err != nil {
			log.Fatalln("Failed to connect/ping to redis", err.Error())
		}
		log.Println("Connected to redis@localhost:6379")
	})
	return client
}
