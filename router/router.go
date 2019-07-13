/*
	路由配置
*/

package router

import (
	"github.com/fishjar/gin-rest-boilerplate/handler"
	"github.com/fishjar/gin-rest-boilerplate/middleware"

	"github.com/gin-gonic/gin"
)

// InitRouter 注入路由，及返回一个gin对象
func InitRouter() *gin.Engine {

	// r := gin.New()
	r := gin.Default() // Default 使用 Logger 和 Recovery 中间件

	r.Use(middleware.BodyLogger()) // 日志中间件
	r.Use(middleware.JWTAuth())    // JWT验证中间件

	r.POST("/account/login", handler.Login) //登录

	// foo模型的路由
	{
		r.POST("/foos", handler.CreateFoo)       // 创建单条
		r.GET("/foos", handler.FindFoos)         // 获取多条
		r.GET("/foos/:id", handler.FindFooByID)  // 按ID查找
		r.PATCH("/foos/:id", handler.UpdateFoo)  // 按ID更新
		r.DELETE("/foos/:id", handler.DeleteFoo) // 按ID删除
		r.POST("/foo", handler.FindOrCreateFoo)  // 查找或创建单条
	}

	return r
}
