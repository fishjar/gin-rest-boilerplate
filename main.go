/*
	GIN + GROM 的 REST 项目模板
*/

package main

import (
	"fmt"

	"github.com/fishjar/gin-rest-boilerplate/config"
	"github.com/fishjar/gin-rest-boilerplate/db"
	"github.com/fishjar/gin-rest-boilerplate/logger"
	"github.com/fishjar/gin-rest-boilerplate/router"
	"github.com/fishjar/gin-rest-boilerplate/utils"
)

func main() {
	defer db.DB.Close()             // 关闭数据库连接
	defer logger.LogFile.Close()    // 关闭日志文件
	defer logger.LogGinFile.Close() // 关闭日志文件
	defer logger.LogReqFile.Close() // 关闭日志文件

	port := config.GetPort()        // 获取端口号
	r := router.InitRouter()        // 获取gin对象
	r.Run(fmt.Sprintf(":%d", port)) // 启动服务

	// 下面是定制启动参数的写法
	// s := &http.Server{
	// 	Addr:           fmt.Sprintf(":%d", config.HTTPPort),
	// 	Handler:        r,
	// 	ReadTimeout:    10 * time.Second,
	// 	WriteTimeout:   10 * time.Second,
	// 	MaxHeaderBytes: 1 << 20,
	// }
	// s.ListenAndServe()
}

func init() {
	env := config.GetEnv()
	if env == "dev" {
		utils.InitDB() //数据初始化
	}
}
