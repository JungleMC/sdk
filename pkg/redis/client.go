package redis

import (
	"github.com/go-redis/redis/v8"
	"log"
)

func NewClient() *redis.Client {
	// TODO: Pull in redis session config from environment variables
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func Close(c *redis.Client) {
	// Ignore nil pointers references, we don't care
	if c == nil {
		return
	}

	if err := c.Close(); err != nil {
		log.Fatalln(err)
	}
}
