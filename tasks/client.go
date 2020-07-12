package tasks

import (
	"github.com/hibiken/asynq"
)

// Client 全局客户端
var Client *asynq.Client

func init() {
	r := asynq.RedisClientOpt{Addr: redisAddr}
	Client = asynq.NewClient(r)
}
