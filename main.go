/*
	GIN + GROM 的 REST 项目模板
*/

package main

import (
	"fmt"
	"os"

	"github.com/fishjar/gin-rest-boilerplate/config"
	"github.com/fishjar/gin-rest-boilerplate/db"
	_ "github.com/fishjar/gin-rest-boilerplate/docs"
	"github.com/fishjar/gin-rest-boilerplate/logger"
	"github.com/fishjar/gin-rest-boilerplate/router"
	"github.com/fishjar/gin-rest-boilerplate/script"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:4000
// @BasePath /
// @query.collection.format multi

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @x-extension-openapi {"example": "value on a json format"}
func main() {
	defer db.DB.Close()             // 关闭数据库连接
	defer db.Redis.Close()          // 关闭Redis连接
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
	fmt.Println("------ GOPATH: ", os.Getenv("GOPATH"))

	// 目录
	if err := os.MkdirAll(config.GetFileDir(), 0755); err != nil {
		panic("上传目录创建失败")
	}

	// 数据
	// TODO 生产环境注意数据
	// env := config.GetEnv()
	// if env == "dev" {
	// 	script.Migrate() // 同步数据表
	// 	script.InitDB()  // 数据初始化
	// }
	script.Migrate() // 同步数据表
	script.InitDB()  // 数据初始化
}
