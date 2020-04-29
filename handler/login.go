/*
	对应路由的处理函数
*/

package handler

import (
	"fmt"
	"net/http"

	"github.com/fishjar/gin-rest-boilerplate/config"
	"github.com/fishjar/gin-rest-boilerplate/db"
	"github.com/fishjar/gin-rest-boilerplate/model"
	"github.com/fishjar/gin-rest-boilerplate/service"
	"github.com/fishjar/gin-rest-boilerplate/utils"
	"github.com/gin-gonic/gin"
)

// LoginAccount 登录处理函数
// TODO：生产环境，错误信息不需要详细情况
// @Summary 帐号登录
// @Description 帐号登录...
// @Tags admin
// @Accept  json
// @Produce  json
// @Param 参数 body model.AuthAccountLoginReq true "登录"
// @Success 200 {object} model.AuthAccountLoginRes
// @Router /admin/login/account [post]
func LoginAccount(c *gin.Context) {

	var loginForm model.AuthAccountLoginReq

	// 绑定数据
	if err := c.ShouldBind(&loginForm); err != nil {
		service.HTTPError(c, "登录失败，参数有误", http.StatusUnauthorized, err)
		return
	}

	// 查询帐号名是否存在
	authType := "account"
	password := utils.MD5Pwd(loginForm.Username, loginForm.Password)
	var auth model.Auth
	if err := db.DB.Where(&model.Auth{
		AuthName: loginForm.Username,
		AuthType: authType,
	}).First(&auth).Error; err != nil {
		service.HTTPError(c, "登录失败，用户名不存在", http.StatusUnauthorized, err)
		return
	}

	// 检查禁用或过期
	if err := auth.CheckEnabled(); err != nil {
		service.HTTPError(c, "登录失败，帐号禁用或过期", http.StatusUnauthorized, err)
		return
	}

	// 验证密码
	if password != *auth.AuthCode {
		service.HTTPError(c, "登录失败，密码错误", http.StatusUnauthorized, fmt.Errorf("username:%s, password:%s", loginForm.Username, loginForm.Password))
		return
	}

	// 查询用户是否存在
	if _, err := service.GetUser(auth.UserID.String()); err != nil {
		service.HTTPError(c, "登录失败，用户数据有误", http.StatusUnauthorized, err)
		return
	}

	// 生成token
	accessToken, err := service.MakeToken(&model.UserJWT{
		AuthID: auth.ID.String(),
		UserID: auth.UserID.String(),
	})
	if err != nil {
		service.HTTPError(c, "登录失败，获取token失败", http.StatusUnauthorized, err)
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
// @Summary 刷新token
// @Description 刷新token...
// @Tags admin
// @Accept  json
// @Produce  json
// @Success 200 {object} model.AuthAccountLoginRes
// @Router /admin/token/refresh [post]
func TokenRefresh(c *gin.Context) {
	AuthID := c.MustGet("AuthID").(string)
	UserID := c.MustGet("UserID").(string)
	newToken, err := service.MakeToken(&model.UserJWT{
		AuthID: AuthID,
		UserID: UserID,
	})
	if err != nil {
		service.HTTPError(c, "刷新token失败", http.StatusInternalServerError, err)
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
