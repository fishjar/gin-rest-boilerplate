/*
	数据结构
*/

package schema

type omit *struct{}

// AccountLoginIn 帐号登录表单
type AccountLoginIn struct {
	UserName string `form:"username" binding:"required"`
	PassWord string `form:"password" binding:"required"`
}

// AccountLoginOut 登录成功返回数据
type AccountLoginOut struct {
	Message     string `json:"msg" binding:"required"`
	TokenType   string `json:"tokenType" binding:"required"`
	AccessToken string `json:"accessToken" binding:"required"`
}

// JWTUser JWT用户数据
type JWTUser struct {
	AuthID string `json:"aid" binding:"required"`
	UserID string `json:"uid" binding:"required"`
}

// PaginQueryIn 分页查询参数
type PaginQueryIn struct {
	Page uint   `form:"page,default=1"`
	Size uint   `form:"size,default=10"`
	Sort string `form:"sort"`
}

// PaginQueryOut 分页查询结果
type PaginQueryOut struct {
	Page  uint        `json:"page" binding:"required"`
	Size  uint        `json:"size" binding:"required"`
	Total uint        `json:"total" binding:"required"`
	Rows  interface{} `json:"rows" binding:"required"`
}
