package model

import (
	"time"

	"github.com/fishjar/gin-rest-boilerplate/db"
	uuid "github.com/satori/go.uuid"
)

// Auth 定义模型
type Auth struct {
	Base
	UserID     uuid.UUID  `json:"userId" gorm:"column:user_id;not null"`                                  // 用户ID
	User       *User      `json:"user" gorm:"foreignkey:UserID"`                                          // 用户
	AuthType   string     `json:"authType" gorm:"column:auth_type;type:VARCHAR(16);not null"`             // 鉴权类型
	AuthName   string     `json:"authName" gorm:"column:auth_name;type:VARCHAR(128);not null"`            // 鉴权名称
	AuthCode   *string    `json:"authCode" gorm:"column:auth_code" binding:"omitempty"`                   // 鉴权识别码
	VerifyTime *time.Time `json:"verifyTime" gorm:"column:verify_time;type:DATETIME" binding:"omitempty"` // 认证时间
	ExpireTime *time.Time `json:"expireTime" gorm:"column:expire_time;type:DATETIME" binding:"omitempty"` // 过期时间
	IsEnabled  *bool      `json:"isEnabled" gorm:"column:is_enabled" binding:"omitempty"`                 // 是否启用
}

// AuthPublic 公开模型
type AuthPublic struct {
	*Auth
	AuthName string  `json:"-" binding:"-"`
	AuthCode *string `json:"-" binding:"-"`
}

// AuthLoginIn 帐号登录表单
type AuthLoginIn struct {
	UserName string `form:"username" binding:"required"`
	PassWord string `form:"password" binding:"required"`
}

// AuthLoginOut 登录成功返回数据
type AuthLoginOut struct {
	Message     string `json:"msg" binding:"required"`
	TokenType   string `json:"tokenType" binding:"required"`
	AccessToken string `json:"accessToken" binding:"required"`
}

// TableName 自定义表名
func (Auth) TableName() string {
	return "auth"
}

// 自定义验证器
// var bookableDate validator.Func = func(fl validator.FieldLevel) bool
// v.RegisterValidation("bookabledate", bookableDate)
// v.RegisterStructValidation(UserStructLevelValidation, User{})

func init() {
	db.DB.AutoMigrate(&Auth{}) // 同步表
}
