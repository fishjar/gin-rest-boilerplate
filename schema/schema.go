/*
	数据结构
*/

package schema

// LoginForm 帐号登录表单
type LoginForm struct {
	UserName string `form:"username" binding:"required"`
	PassWord string `form:"password" binding:"required"`
}

// LoginRes 登录成功返回数据
type LoginRes struct {
	Message     string `json:"message" binding:"required"`
	TokenType   string `json:"tokenType" binding:"required"`
	AccessToken string `json:"accessToken" binding:"required"`
}

// JWTUser JWT表单
type JWTUser struct {
	AuthID string `json:"aid" binding:"required"`
	UserID string `json:"uid" binding:"required"`
}
