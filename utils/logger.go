package utils

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/fishjar/gin-rest-boilerplate/config"

	"github.com/gin-gonic/gin"
)

// 定义日志记录器
var (
	LogFile    *os.File
	LogDebug   *log.Logger
	LogInfo    *log.Logger
	LogWarning *log.Logger
	LogError   *log.Logger
)

func init() {

	// 测试环境
	logPath := os.Getenv("GOPATH") + config.LogPath
	LogFile, err := os.Create(logPath)

	// 生产环境不必每次都新建文件，打开即可
	// LogFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic("创建日志文件失败")
	}
	fmt.Println("日志文件：", logPath)

	LogDebug = log.New(LogFile, "[Debug]", log.LstdFlags)
	LogInfo = log.New(LogFile, "[Info]", log.LstdFlags)
	LogWarning = log.New(LogFile, "[Warning]", log.LstdFlags)
	LogError = log.New(LogFile, "[Error]", log.LstdFlags)

	// 配置gin的log
	// 测试环境
	gin.DefaultWriter = io.MultiWriter(LogFile, os.Stdout)

	// 生产环境不需要打印到命令行
	// gin.DefaultWriter = io.MultiWriter(LogFile)
}
