/*
	中间件
*/

package middleware

import (
	"net/http"
	"strings"

	"github.com/fishjar/gin-rest-boilerplate/utils"
	"github.com/gin-gonic/gin"
)

// JWTAuth 验证中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 登录链接不做验证
		path := c.Request.URL.Path
		if path == "/account/login" {
			c.Next()
			return
		}

		// 获取token
		authorization := c.Request.Header.Get("Authorization")
		utils.LogError.Println("req authorization: ", authorization)
		tokenString := strings.Replace(authorization, "Bearer ", "", 1)

		if tokenString != "" {
			if claims, err := utils.ParseToken(tokenString); claims != nil && err == nil {

				// 验证成功
				// 挂载到全局
				c.Set("userName", claims.UserName)
				c.Set("userType", claims.UserType)

				// 返回一个新token给客户端(未验证)
				if newToken, err := utils.MakeToken(claims.UserName, claims.UserType); err == nil {
					c.Writer.Header().Set("authtoken", newToken)
				}

				c.Next()
				return
			}
		}

		// 验证失败
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "没有权限",
		})
		c.Abort() // 直接返回
	}
}
