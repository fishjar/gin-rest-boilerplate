package crons

import "github.com/robfig/cron"

// 定时任务
const (
	EVERY1SENCOND = "@every 1s" // 每1秒
	EVERY2SENCOND = "@every 2s" // 每2秒
	EVERY3SENCOND = "@every 3s" // 每3秒
)

// Cron 定时任务
var Cron *cron.Cron

func init() {
	c := cron.New()
	c.AddJob(EVERY1SENCOND, &TestJob{"1s"})
	c.AddFunc(EVERY2SENCOND, TestFunc("2s"))
	c.Start()
	Cron = c
}
