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
	r := gin.Default()                    // Default 使用 Logger 和 Recovery 中间件
	r.Use(middleware.LoggerToFile())      // 日志中间件
	r.GET("/ping", func(c *gin.Context) { // pingpong
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/login/account", handler.LoginAccount) //登录

	authorized := r.Group("/admin")      // JWT验证路由组
	authorized.Use(middleware.JWTAuth()) // JWT验证中间件
	{

		authorized.GET("/auths", handler.AuthFindAndCountAll)    // 获取多条
		authorized.GET("/auths/:id", handler.AuthFindByPk)       // 按ID查找
		authorized.POST("/auths", handler.AuthSingleCreate)      // 创建单条
		authorized.PATCH("/auths/:id", handler.AuthUpdateByPk)   // 按ID更新
		authorized.DELETE("/auths/:id", handler.AuthDestroyByPk) // 按ID删除
		authorized.POST("/auth", handler.AuthFindOrCreate)       // 查询或创建
		authorized.PATCH("/auths", handler.AuthUpdateBulk)       // 批量更新
		authorized.DELETE("/auths", handler.AuthDestroyBulk)     // 批量删除
	}
	{
		authorized.GET("/groups", handler.GroupFindAndCountAll)    // 获取多条
		authorized.GET("/groups/:id", handler.GroupFindByPk)       // 按ID查找
		authorized.POST("/groups", handler.GroupSingleCreate)      // 创建单条
		authorized.PATCH("/groups/:id", handler.GroupUpdateByPk)   // 按ID更新
		authorized.DELETE("/groups/:id", handler.GroupDestroyByPk) // 按ID删除
		authorized.POST("/group", handler.GroupFindOrCreate)       // 查询或创建
		authorized.PATCH("/groups", handler.GroupUpdateBulk)       // 批量更新
		authorized.DELETE("/groups", handler.GroupDestroyBulk)     // 批量删除
	}
	{
		authorized.GET("/menus", handler.MenuFindAndCountAll)    // 获取多条
		authorized.GET("/menus/:id", handler.MenuFindByPk)       // 按ID查找
		authorized.POST("/menus", handler.MenuSingleCreate)      // 创建单条
		authorized.PATCH("/menus/:id", handler.MenuUpdateByPk)   // 按ID更新
		authorized.DELETE("/menus/:id", handler.MenuDestroyByPk) // 按ID删除
		authorized.POST("/menu", handler.MenuFindOrCreate)       // 查询或创建
		authorized.PATCH("/menus", handler.MenuUpdateBulk)       // 批量更新
		authorized.DELETE("/menus", handler.MenuDestroyBulk)     // 批量删除
	}
	{
		authorized.GET("/roles", handler.RoleFindAndCountAll)    // 获取多条
		authorized.GET("/roles/:id", handler.RoleFindByPk)       // 按ID查找
		authorized.POST("/roles", handler.RoleSingleCreate)      // 创建单条
		authorized.PATCH("/roles/:id", handler.RoleUpdateByPk)   // 按ID更新
		authorized.DELETE("/roles/:id", handler.RoleDestroyByPk) // 按ID删除
		authorized.POST("/role", handler.RoleFindOrCreate)       // 查询或创建
		authorized.PATCH("/roles", handler.RoleUpdateBulk)       // 批量更新
		authorized.DELETE("/roles", handler.RoleDestroyBulk)     // 批量删除
	}
	{
		authorized.GET("/users", handler.UserFindAndCountAll)    // 获取多条
		authorized.GET("/users/:id", handler.UserFindByPk)       // 按ID查找
		authorized.POST("/users", handler.UserSingleCreate)      // 创建单条
		authorized.PATCH("/users/:id", handler.UserUpdateByPk)   // 按ID更新
		authorized.DELETE("/users/:id", handler.UserDestroyByPk) // 按ID删除
		authorized.POST("/user", handler.UserFindOrCreate)       // 查询或创建
		authorized.PATCH("/users", handler.UserUpdateBulk)       // 批量更新
		authorized.DELETE("/users", handler.UserDestroyBulk)     // 批量删除
		authorized.GET("/user/roles", handler.UserFindMyRoles)   // 获取角色列表
		authorized.GET("/user/menus", handler.UserFindMyMenus)   // 获取菜单列表
	}
	{
		authorized.GET("/usergroups", handler.UserGroupFindAndCountAll)    // 获取多条
		authorized.GET("/usergroups/:id", handler.UserGroupFindByPk)       // 按ID查找
		authorized.POST("/usergroups", handler.UserGroupSingleCreate)      // 创建单条
		authorized.PATCH("/usergroups/:id", handler.UserGroupUpdateByPk)   // 按ID更新
		authorized.DELETE("/usergroups/:id", handler.UserGroupDestroyByPk) // 按ID删除
		authorized.POST("/usergroup", handler.UserGroupFindOrCreate)       // 查询或创建
		authorized.PATCH("/usergroups", handler.UserGroupUpdateBulk)       // 批量更新
		authorized.DELETE("/usergroups", handler.UserGroupDestroyBulk)     // 批量删除
	}

	return r
}
