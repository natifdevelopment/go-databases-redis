package redis

import (
	"context"
	"fmt"
	"log"

	configs "github.com/natifdevelopment/go-config"
	"github.com/redis/go-redis/v9"
)

var RedisConn *redis.Client

func SetupRedis() {
	addr := fmt.Sprintf("%s:%s", configs.REDIS_HOST, configs.REDIS_PORT)
	RedisConn = redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: configs.REDIS_USERNAME,
		Password: configs.REDIS_PASSWORD,
		DB:       configs.REDIS_DB,
	})

	ctx := context.Background()
	if err := RedisConn.Ping(ctx).Err(); err != nil {
		log.Printf("Warning: Redis connection failed: %v", err)
		RedisConn = nil
		return
	}
	log.Println("Redis connection established successfully")
}

func CloseRedis() {
	if RedisConn != nil {
		RedisConn.Close()
	}
}
