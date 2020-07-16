package tasks

import (
	"github.com/fishjar/gin-rest-boilerplate/config"
	"github.com/hibiken/asynq"
)

// Client 全局客户端
var Client *asynq.Client

func init() {
	r := asynq.RedisClientOpt{Addr: config.RedisAddr}
	Client = asynq.NewClient(r)
}
