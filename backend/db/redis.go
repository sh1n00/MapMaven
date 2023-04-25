package db

import (
	"github.com/go-redis/redis"
	"log"
)

type RedisHandler struct {
	Conn *redis.Client
}

var RedisClient *RedisHandler = NewRedisHandler()

func NewRedisHandler() *RedisHandler {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	if _, err := redisClient.Ping().Result(); err != nil {
		log.Fatalln(err)
	}
	return &RedisHandler{Conn: redisClient}
}
