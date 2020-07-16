package db

import (
	"fmt"

	"github.com/fishjar/gin-rest-boilerplate/config"
	"github.com/go-redis/redis/v7"
)

// Redis 全局实例
var Redis *redis.Client

func init() {
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr, // redis 地址
		Password: config.RedisPWD,  // redis 密码
		DB:       0,                // use default DB
	})

	if pong, err := client.Ping().Result(); err != nil {
		fmt.Println("----redis ping----", pong, err)
		panic("redis连接错误")
	}

	Redis = client
}
