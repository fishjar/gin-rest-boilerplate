/*
	中间件
*/

package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/fishjar/gin-rest-boilerplate/logger"
	"github.com/fishjar/gin-rest-boilerplate/model"
	"github.com/fishjar/gin-rest-boilerplate/service"

	"github.com/fishjar/gin-rest-boilerplate/utils"
	"github.com/gin-gonic/gin"
)

// JWTAuth 验证中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 获取token
		// authorization := c.Request.Header.Get("Authorization")
		authorization := c.GetHeader("Authorization")
		accessToken := strings.Replace(authorization, "Bearer ", "", 1)

		// token 为空
		if len(accessToken) == 0 {
			// 验证失败
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.HTTPError{
				Code:    401,
				Message: "没有权限：token不能为空!",
			})
			return
		}

		// 解析token
		claims, err := utils.ParseToken(accessToken)
		if claims == nil || err != nil {
			// 验证失败
			go logger.Log.WithFields(logrus.Fields{
				"accessToken": accessToken,
			}).Warn("JWTAuth 认证失败")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "没有权限：JWT验证失败",
			})
			return
		}

		AuthID := claims.Subject
		UserID := claims.UserID

		// 从数据库验证 AuthID 和 UserID 有效性
		auth, err := service.GetAuthWithRoles(AuthID)
		if err != nil { // 帐号不存在
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "没有权限：帐号不存在",
			})
			return
		}
		if err := auth.CheckEnabled(); err != nil { // 禁用或过期
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "没有权限：禁用或过期",
			})
			return
		}
		if auth.User.ID.String() != UserID { // 用户数据有误
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "没有权限：用户数据有误",
			})
			return
		}
		// 获取用户角色列表
		roles := make([]string, len(auth.User.Roles))
		for _, role := range auth.User.Roles {
			roles = append(roles, role.Name)
		}

		// TODO: 从redis认证 AuthID 和 UserID 有效性

		// 验证成功
		// 挂载到全局
		fmt.Println("----token auth info----")
		fmt.Println(AuthID)
		fmt.Println(UserID)
		fmt.Println(roles)
		c.Set("AuthID", AuthID)
		c.Set("UserID", UserID)
		c.Set("UserRoles", roles)

		// 返回一个新token给客户端(暂不需要)
		// if newToken, err := utils.MakeToken(&model.UserJWT{
		// 	AuthID: AuthID,
		// 	UserID: UserID,
		// }); err == nil {
		// 	c.Writer.Header().Set("X-Authorization", newToken)
		// }

		c.Next()

	}
}
