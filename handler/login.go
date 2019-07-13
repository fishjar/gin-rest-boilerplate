/*
	对应路由的处理函数
*/

package handler

import (
	"net/http"

	"github.com/fishjar/gin-rest-boilerplate/db"
	"github.com/fishjar/gin-rest-boilerplate/model"

	"github.com/fishjar/gin-rest-boilerplate/utils"
	"github.com/gin-gonic/gin"
)

// LoginForm 登录表单
type LoginForm struct {
	UserName string `form:"userName" binding:"required"`
	UserType string `form:"userType" binding:"required"`
	PassWord string `form:"password" binding:"required"`
}

// Login 登录处理函数
func Login(c *gin.Context) {

	var login LoginForm

	// 绑定数据
	if err := c.ShouldBind(&login); err != nil {
		utils.LogWarning.Println("登录失败，参数有误：", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "登录失败",
		})
		return
	}

	// 查询用户是否存在
	var user model.User
	if err := db.DB.Where(&model.User{
		UserName: login.UserName,
		UserType: login.UserType,
		PassWord: utils.MD5Pwd(login.PassWord),
	}).First(&user).Error; err != nil {
		utils.LogWarning.Println("登录失败，用户名或密码错误：", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "登录失败",
		})
		return
	}

	// 生成token
	authtoken, err := utils.MakeToken(user.UserName, user.UserType)
	if err != nil {
		utils.LogWarning.Println("登录失败，获取token失败：", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "登录失败",
		})
		return
	}

	// 登录成功
	c.JSON(http.StatusOK, gin.H{
		"message":   "登录成功",
		"authtoken": authtoken,
	})

}
