package handler

import (
	"fmt"
	"time"

	"github.com/bsm/redislock"
	"github.com/fishjar/gin-rest-boilerplate/crons"
	"github.com/fishjar/gin-rest-boilerplate/locker"
	"github.com/fishjar/gin-rest-boilerplate/tasks"
	"github.com/gin-gonic/gin"
)

// Pong 测试handle
func Pong(c *gin.Context) {
	// 获取锁
	lock, err := locker.Locker.Obtain(locker.PING, 10*1000*time.Millisecond, nil)
	if err == redislock.ErrNotObtained {
		// fmt.Println("Could not obtain lock!")
		c.JSON(200, gin.H{
			"message": "Could not obtain lock!",
		})
		return
	} else if err != nil {
		// log.Fatalln(err)
		c.JSON(200, gin.H{
			"message": "obtain lock err",
		})
		return
	}
	// 释放锁
	defer lock.Release()

	// 测试创建任务队列
	t := tasks.NewEmailDeliveryTask(42, "some:template:id") // 创建任务
	if _, err := tasks.Client.Enqueue(t); err != nil {      // 添加到任务队列
		fmt.Println("添加任务队列失败", err)
		c.JSON(200, gin.H{
			"message": "添加任务队列失败",
		})
		return
	}

	// 测试创建定时任务
	crons.Cron.AddJob(crons.EVERY3SENCOND, &crons.TestJob{Value: "3s"})

	time.Sleep(5 * 1000 * time.Millisecond)

	c.JSON(200, gin.H{
		"message": "pong..",
	})
}
