/*
	GIN + GROM 的 REST 项目模板
*/

package main

import (
	"fmt"

	"github.com/fishjar/gin-rest-boilerplate/config"
	"github.com/fishjar/gin-rest-boilerplate/db"
	"github.com/fishjar/gin-rest-boilerplate/router"
	"github.com/fishjar/gin-rest-boilerplate/utils"
)

func main() {
	defer db.DB.Close()         // 关闭数据库连接
	defer utils.LogFile.Close() // 关闭日志文件

	r := router.InitRouter()                   // 获取gin对象
	r.Run(fmt.Sprintf(":%d", config.HTTPPort)) // 启动服务

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
