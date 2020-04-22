package model

import (
	"github.com/fishjar/gin-rest-boilerplate/db"
)

// Role 定义模型
type Role struct {
	Base
	Name  string  `json:"name" gorm:"column:name;type:VARCHAR(32);unique;not null" binding:"min=3,max=20"` // 角色名称
	Users []*User `json:"users" gorm:"many2many:userrole;"`                                                // 用户
	Menus []*Menu `json:"menus" gorm:"many2many:rolemenu;"`                                                // 菜单
}

// TableName 自定义表名
func (Role) TableName() string {
	return "role"
}

func init() {
	db.DB.AutoMigrate(&Role{}) // 同步表
}
