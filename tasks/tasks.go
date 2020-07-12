package tasks

import (
	"github.com/fishjar/gin-rest-boilerplate/config"
)

var redisAddr = config.GetRedisAddr()

// 任务列表
const (
	emailDelivery   = "email:deliver"
	imageProcessing = "image:process"
)
