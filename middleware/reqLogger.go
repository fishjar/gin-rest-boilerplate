package middleware

import (
	"time"

	"github.com/fishjar/gin-rest-boilerplate/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// LoggerToFile 日志记录到文件
func LoggerToFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()                    // 开始时间
		reqMethod := c.Request.Method              // 请求方式
		reqURI := c.Request.RequestURI             // 请求路由
		clientIP := c.ClientIP()                   // 请求IP
		c.Next()                                   // 处理请求
		statusCode := c.Writer.Status()            // 返回状态码
		endTime := time.Now()                      // 结束时间
		latencyTime := endTime.Sub(startTime)      // 执行时间
		go logger.LogReq.WithFields(logrus.Fields{ // 日志记录
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqURI,
		}).Info()
	}
}

// LoggerToMongo 日志记录到 MongoDB
func LoggerToMongo() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

// LoggerToES 日志记录到 ES
func LoggerToES() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

// LoggerToMQ 日志记录到 MQ
func LoggerToMQ() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
