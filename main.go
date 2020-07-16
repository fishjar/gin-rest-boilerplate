/*
	GIN + GROM 的 REST 项目模板
*/

package main

import (
	"fmt"
	"os"

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

	taskDone := make(chan bool, 1)
	allDone := make(chan bool, 1)

	go router.RunGinServer(taskDone, allDone) // 启动gin服务
	go router.RunTaskServer(taskDone)         // 启动task服务

	<-allDone // 阻塞，等待全部退出
	fmt.Println("------all exited--------")
}

func init() {
	fmt.Println("------ GOPATH----------")
	fmt.Println(os.Getenv("GOPATH"))

	// 数据
	// TODO 生产环境注意数据
	// env := config.GINENV
	// if env == "dev" {
	// 	script.Migrate() // 同步数据表
	// 	script.InitDB()  // 数据初始化
	// }
	script.Migrate() // 同步数据表
	script.InitDB()  // 数据初始化
}
