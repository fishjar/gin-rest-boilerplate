/*
	中间件
*/

package middleware

import (
	"net/http"
	"strings"

	"github.com/fishjar/gin-rest-boilerplate/schema"
	"github.com/sirupsen/logrus"

	"github.com/fishjar/gin-rest-boilerplate/logger"

	"github.com/fishjar/gin-rest-boilerplate/utils"
	"github.com/gin-gonic/gin"
)

// JWTAuth 验证中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 获取token
		authorization := c.Request.Header.Get("Authorization")
		accessToken := strings.Replace(authorization, "Bearer ", "", 1)

		// token 为空
		if accessToken == "" {
			// 验证失败
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"msg": "没有权限：token不能为空",
			})
			return
		}

		// 解析token
		claims, err := utils.ParseToken(accessToken)
		if claims == nil || err != nil {
			// 验证失败
			logger.Log.WithFields(logrus.Fields{
				"accessToken": accessToken,
			}).Warn("JWTAuth 认证失败")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"msg": "没有权限：token解析错误",
			})
			return
		}

		AuthID := claims.AuthID
		UserID := claims.Subject

		// TODO: 从redis认证 AuthID 和 UserID 有效性

		// var auth model.Auth
		// // 验证帐号
		// if err := db.DB.Where("id = ?", authID).First(&auth).Error; err != nil {
		// 	// 验证失败
		// 	c.JSON(http.StatusUnauthorized, gin.H{
		// 		"code":    401,
		// 		"msg": "没有权限：帐号不存在",
		// 	})
		// 	c.Abort() // 直接返回
		// 	return
		// }
		// // 验证密码
		// if auth.AuthName != authName || auth.AuthType != authType {
		// 	// 验证失败
		// 	c.JSON(http.StatusUnauthorized, gin.H{
		// 		"code":    401,
		// 		"msg": "没有权限：帐号或密码错误",
		// 	})
		// 	c.Abort() // 直接返回
		// 	return
		// }

		// 验证成功
		// 挂载到全局
		// c.Set("AuthID", AuthID)
		// c.Set("UserID", UserID)
		c.Set("user", schema.JWTUser{
			AuthID: AuthID,
			UserID: UserID,
		})

		// 返回一个新token给客户端(未验证)
		if newToken, err := utils.MakeToken(&schema.JWTUser{
			AuthID: AuthID,
			UserID: UserID,
		}); err == nil {
			c.Writer.Header().Set("X-Authorization", newToken)
		}

		c.Next()

	}
}
