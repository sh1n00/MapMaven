package db

import (
	"backend/utils"
	"encoding/json"
	"github.com/go-redis/redis"
	"log"
)

type RedisHandler struct {
	Conn *redis.Client
}

var RedisClient = NewRedisHandler()

func init() {
	if err := RedisClient.Conn.FlushDB().Err(); err != nil {
		log.Fatalln(err)
	}

	instructions := []string{
		"情報科学B棟3階にサイバネティクス・リアリティ工学研究室があります。",
		"サイバネティクスリアリティ工学研究室はVRやHCIをテーマとした研究を行っています。",
		"情報科学A棟7階にインタラクティブメディアデザイン研究室があります。",
		"インタラクティブメディアデザイン研究室はVRやARやCVをテーマとした研究を行っています。",
		"情報科学A棟6階に自然言語処理研究室があります。",
		"自然言語処理学研究室は自然言語を中心とした研究を行っています。",
		"A棟のエレベーターは、この場所から右側奥にあります。",
		"B棟のエレベーターは、この場所から左側にあります。",
	}

	for _, instruction := range instructions {
		embedding, err := utils.Embeddings(instruction)
		if err != nil {
			log.Println(err)
		}
		embeddingBytes, err := json.Marshal(embedding)
		if err != nil {
			log.Println(err)
		}
		if err = RedisClient.Conn.Set(instruction, embeddingBytes, 0).Err(); err != nil {
			log.Println(err)
		}
	}
}

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
