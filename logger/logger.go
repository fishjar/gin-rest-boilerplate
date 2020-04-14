/*
	日志记录器
*/

package logger

import (
	"io"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// LogFile 日志文件
var LogFile *os.File

// LogGinFile 日志文件
var LogGinFile *os.File

// LogReqFile 日志文件
var LogReqFile *os.File

// Log 日志记录器
var Log = logrus.New()

// LogReq 日志记录器
var LogReq = logrus.New()

func init() {
	// 获取日志路径
	rootDir, _ := os.Getwd()
	logDir := path.Join(rootDir, "tmp/log")

	// 创建日志路径
	err := os.MkdirAll(logDir, 0755)
	if err != nil {
		panic("创建日志目录失败")
	}

	// 创建日志文件
	LogFile, err := os.Create(path.Join(logDir, "log.log"))
	if err != nil {
		panic("创建日志文件失败")
	}
	// 创建GIN日志文件
	LogGinFile, err := os.Create(path.Join(logDir, "gin.log"))
	if err != nil {
		panic("创建GIN日志文件失败")
	}
	// 创建REQ日志文件
	LogReqFile, err := os.Create(path.Join(logDir, "req.log"))
	if err != nil {
		panic("创建REQ日志文件失败")
	}

	// 配置日志记录器文件
	Log.Out = LogFile
	LogReq.Out = LogReqFile

	// 配置GIN日志文件
	// gin.DefaultWriter = io.MultiWriter(LogGinFile)
	gin.DefaultWriter = io.MultiWriter(LogGinFile, os.Stdout)

}
