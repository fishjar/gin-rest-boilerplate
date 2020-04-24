/*
	对应路由的处理函数
*/

package handler

import (
	"net/http"

	"github.com/fishjar/gin-rest-boilerplate/config"
	"github.com/fishjar/gin-rest-boilerplate/db"
	"github.com/fishjar/gin-rest-boilerplate/logger"
	"github.com/fishjar/gin-rest-boilerplate/model"
	"github.com/fishjar/gin-rest-boilerplate/service"
	"github.com/fishjar/gin-rest-boilerplate/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// LoginAccount 登录处理函数
// TODO：生产环境，错误信息不需要详细情况
func LoginAccount(c *gin.Context) {

	var loginForm model.AuthAccountLoginReq

	// 绑定数据
	if err := c.ShouldBind(&loginForm); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "登录失败，参数有误",
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
			"msg": "登录失败，用户名或密码错误",
		})
		return
	}

	// 检查禁用或过期
	if err := service.AuthCheck(auth); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "登录失败，帐号禁用或过期",
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
			"msg": "登录失败，用户名或密码错误",
		})
		return
	}

	// 查询用户是否存在
	if _, err := service.GetUser(auth.UserID.String()); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "登录失败，用户数据有误",
		})
		return
	}

	// 生成token
	accessToken, err := utils.MakeToken(&model.UserJWT{
		AuthID: auth.ID.String(),
		UserID: auth.UserID.String(),
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "登录失败，获取token失败",
		})
		return
	}

	// TODO：保存到redis

	// 登录成功
	c.JSON(http.StatusOK, model.AuthAccountLoginRes{
		Message:     "登录成功",
		TokenType:   "bearer",
		AccessToken: accessToken,
		ExpiresIn:   config.JWTExpiresIn,
	})

}

// TokenRefresh 刷新token
func TokenRefresh(c *gin.Context) {
	AuthID := c.MustGet("AuthID").(string)
	UserID := c.MustGet("UserID").(string)
	newToken, err := utils.MakeToken(&model.UserJWT{
		AuthID: AuthID,
		UserID: UserID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "刷新token失败",
		})
		return
	}

	// TODO：保存到redis

	// 更新成功
	c.JSON(http.StatusOK, model.AuthAccountLoginRes{
		Message:     "刷新成功",
		TokenType:   "bearer",
		AccessToken: newToken,
		ExpiresIn:   config.JWTExpiresIn,
	})
}
