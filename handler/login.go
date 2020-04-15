/*
	对应路由的处理函数
*/

package handler

import (
	"net/http"

	"github.com/fishjar/gin-rest-boilerplate/db"
	"github.com/fishjar/gin-rest-boilerplate/logger"
	"github.com/fishjar/gin-rest-boilerplate/model"
	"github.com/fishjar/gin-rest-boilerplate/schema"
	"github.com/fishjar/gin-rest-boilerplate/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// LoginAccount 登录处理函数
func LoginAccount(c *gin.Context) {

	var loginForm schema.LoginForm

	// 绑定数据
	if err := c.ShouldBind(&loginForm); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "登录失败，参数有误",
		})
		return
	}

	// 查询帐号名是否存在
	authType := "account"
	passWord := utils.MD5Pwd(loginForm.UserName, loginForm.PassWord)
	var auth model.Auth
	if err := db.DB.Where(&model.Auth{
		AuthName: loginForm.UserName,
		AuthType: authType,
	}).First(&auth).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "登录失败，用户名或密码错误",
		})
		return
	}

	// 验证密码
	if passWord != *auth.AuthCode {
		logger.Log.WithFields(logrus.Fields{
			"username": loginForm.UserName,
			"password": loginForm.PassWord,
		}).Warn("登录失败，密码错误")
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "登录失败，用户名或密码错误",
		})
		return
	}

	// 查询用户是否存在
	var user model.User
	if err := db.DB.Where("id = ?", auth.UserID).Preload("Roles").First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "登录失败，用户数据有误",
		})
		return
	}

	// 生成token
	id := auth.ID.String()
	authtoken, err := utils.MakeToken(id, auth.AuthName, auth.AuthType)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "登录失败，获取token失败",
		})
		return
	}

	// 登录成功
	c.JSON(http.StatusOK, gin.H{
		"message":     "登录成功",
		"tokenType":   "bearer",
		"accessToken": authtoken,
	})

}
