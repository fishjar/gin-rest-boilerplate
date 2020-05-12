/*
	中间件
*/

package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/fishjar/gin-rest-boilerplate/db"
	"github.com/fishjar/gin-rest-boilerplate/model"
	"github.com/fishjar/gin-rest-boilerplate/service"

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
			service.HTTPAbortError(c, "没有权限：token不能为空", http.StatusUnauthorized, nil)
			return
		}

		// 解析token
		claims, err := service.ParseToken(accessToken)
		if claims == nil || err != nil {
			// 验证失败
			service.HTTPAbortError(c, "没有权限：JWT验证失败", http.StatusUnauthorized, fmt.Errorf("token:%s", accessToken))
			return
		}

		aid := claims.Subject
		uid := claims.UserID

		// // 从数据库验证 aid 和 uid 有效性
		// auth, err := service.GetAuthWithUser(aid)
		// if err != nil { // 帐号不存在
		// 	service.HTTPAbortError(c, "没有权限：帐号不存在", http.StatusUnauthorized, fmt.Errorf("token:%s", accessToken))
		// 	return
		// }
		// if err := auth.CheckEnabled(); err != nil { // 禁用或过期
		// 	service.HTTPAbortError(c, "没有权限：禁用或过期", http.StatusUnauthorized, fmt.Errorf("token:%s", accessToken))
		// 	return
		// }
		// if auth.User.ID.String() != uid { // 用户数据有误
		// 	service.HTTPAbortError(c, "没有权限：用户数据有误", http.StatusUnauthorized, fmt.Errorf("token:%s", accessToken))
		// 	return
		// }
		// roles, err := auth.User.GetRoles() // 获取用户角色列表
		// if err != nil {
		// 	service.HTTPAbortError(c, "没有权限：帐号角色信息有误", http.StatusUnauthorized, fmt.Errorf("token:%s", accessToken))
		// 	return
		// }
		// roleNames := service.RolesToNames(roles)

		// 从redis认证 aid 和 uid 有效性
		userInfo, err := db.Redis.HGetAll("user:" + uid).Result()
		roleNames := strings.Split(userInfo["roles"], ",")
		if err != nil {
			service.HTTPAbortError(c, "没有权限：缓存信息不存在", http.StatusUnauthorized, fmt.Errorf("token:%s", accessToken))
			return
		}
		if userInfo["aid"] != aid {
			service.HTTPAbortError(c, "没有权限：帐号ID不一致", http.StatusUnauthorized, fmt.Errorf("token:%s", accessToken))
			return
		}
		// TODO: 校验JWT签发时间

		// 验证成功
		// 当前用户信息挂载到内存
		c.Set("UserInfo", model.UserCurrent{
			AuthID: aid,
			UserID: uid,
			Roles:  roleNames,
		})

		// 返回一个新token给客户端(每次自动刷新token)
		// if newToken, err := service.MakeToken(&model.UserJWT{
		// 	aid: aid,
		// 	uid: uid,
		// }); err == nil {
		// 	c.Writer.Header().Set("X-Authorization", newToken)
		// }

		c.Next()

	}
}
