package model

import (
	"github.com/fishjar/gin-rest-boilerplate/db"
	uuid "github.com/satori/go.uuid"
)

// Menu 定义模型
type Menu struct {
	Base
	ParentID uuid.UUID `json:"parentId" gorm:"column:parent_id" binding:"omitempty"` // 父ID
	Parent   *Menu     `json:"parent" gorm:"foreignkey:ParentID"`                    // 父菜单
	Children []*Menu   `json:"children" gorm:"foreignkey:ParentID"`                  // 子菜单
	Name     string    `json:"name" gorm:"column:name;not null"`                     // 菜单名称
	Path     string    `json:"path" gorm:"column:path;not null"`                     // 菜单路径
	Icon     *string   `json:"icon" gorm:"column:icon" binding:"omitempty"`          // 菜单图标
	Sort     *int      `json:"sort" gorm:"column:sort" binding:"omitempty"`          // 排序
	Roles    []*Role   `json:"roles" gorm:"many2many:rolemenu;"`                     // 角色
}

// TableName 自定义表名
func (Menu) TableName() string {
	return "menu"
}

func init() {
	db.DB.AutoMigrate(&Menu{}) // 同步表
}
