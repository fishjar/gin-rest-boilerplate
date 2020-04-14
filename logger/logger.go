package logger

import (
	"os"
	"path"

	"github.com/fishjar/gin-rest-boilerplate/config"

	"github.com/sirupsen/logrus"
)

// LogFile 日志文件
var LogFile *os.File

// Log 日志
var Log *logrus.Logger = logrus.New()

func init() {
	// 获取日志路径
	rootPath, _ := os.Getwd()
	logPath := path.Join(rootPath, "tmp/log")

	// 创建日志路径
	err := os.MkdirAll(logPath, 0755)
	if err != nil {
		panic("创建日志目录失败")
	}

	// 创建日志文件
	LogFile, err := os.Create(path.Join(logPath, config.LogName))
	// defer LogFile.Close()
	if err != nil {
		panic("创建日志文件失败")
	}

	// Logger := logrus.New()
	// log.Out = os.Stdout
	Log.Out = LogFile
	// Log := log.WithFields(logrus.Fields{})

	// Log.Info("I'll be logged with common and other field")
	// Log.Warn("Me too")
	Log.Warn("test warn")
}
