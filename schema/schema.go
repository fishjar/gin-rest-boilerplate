/*
	数据结构
*/

package schema

// LoginForm 登录表单
type LoginForm struct {
	UserName string `form:"username" binding:"required"`
	PassWord string `form:"password" binding:"required"`
}

// JWTForm JWT表单
type JWTForm struct {
	AuthID   string `form:"username" binding:"required"`
	AuthName string `form:"username" binding:"required"`
	AuthType string `form:"password" binding:"required"`
}
