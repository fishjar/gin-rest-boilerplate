/*
	GIN + GROM 的 REST 项目模板
*/

package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fishjar/gin-rest-boilerplate/config"
	"github.com/fishjar/gin-rest-boilerplate/db"
	_ "github.com/fishjar/gin-rest-boilerplate/docs"
	"github.com/fishjar/gin-rest-boilerplate/logger"
	"github.com/fishjar/gin-rest-boilerplate/router"
	"github.com/fishjar/gin-rest-boilerplate/script"
	"github.com/fishjar/gin-rest-boilerplate/tasks"
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
	defer tasks.Client.Close()      // 关闭任务队列服务
	defer db.DB.Close()             // 关闭数据库连接
	defer db.Redis.Close()          // 关闭Redis连接
	defer logger.LogFile.Close()    // 关闭日志文件
	defer logger.LogGinFile.Close() // 关闭日志文件
	defer logger.LogReqFile.Close() // 关闭日志文件
	runServer()                     // 启动服务
}

func runServer() {
	done := make(chan bool, 1)
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

	r := router.InitRouter() // 获取gin对象
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.HTTPPort),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() { // gin服务
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("gin revocer")
			}
		}()

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed { //阻塞，等待关闭或错误
			fmt.Println(err)
			panic("gin启动失败")
		}
	}()

	go func() { // 任务队列服务
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("tasks sever revocer")
			}
		}()

		if err := tasks.Server(); err != nil { // 阻塞，等待退出信号
			fmt.Println(err)
			panic("任务队列启动失败")
		}
		done <- true
	}()

	sig := <-sigs // 阻塞，等待退出信号
	fmt.Println("--------got a signal-------", sig)

	<-done // 阻塞，等待任务队列安全退出
	fmt.Println("任务队列服务已退出")

	now := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil { // 退出gin服务
		fmt.Println("关闭gin服务失败", err)
	}

	select {
	case <-ctx.Done(): // 阻塞，等待3秒
		fmt.Println("----timeout of 3 seconds-----")
	}

	fmt.Println("------exited--------", time.Since(now))
}

func init() {
	fmt.Println("------ GOPATH----------")
	fmt.Println(os.Getenv("GOPATH"))

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
