package db

import (
	"fmt"

	"github.com/go-redis/redis/v7"
)

// Redis 全局实例
var Redis *redis.Client

func init() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if pong, err := client.Ping().Result(); err != nil {
		fmt.Println("----redis ping----", pong, err)
		panic("redis连接错误")
	}

	Redis = client
}
