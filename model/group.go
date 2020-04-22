package model

import (
	"github.com/fishjar/gin-rest-boilerplate/db"
	uuid "github.com/satori/go.uuid"
)

// Group 定义模型
type Group struct {
	Base
	Name     string    `json:"name" gorm:"column:name;type:VARCHAR(32);not null" binding:"min=3,max=20"` // 组名称
	LeaderID uuid.UUID `json:"leaderId" gorm:"column:leader_id;not null"`                                // 队长ID
	Leader   *User     `json:"leader" gorm:"foreignkey:LeaderID"`                                        // 队长
	Users    []*User   `json:"users" gorm:"many2many:usergroup;"`                                        // 队员
}

// TableName 自定义表名
func (Group) TableName() string {
	return "group"
}

func init() {
	db.DB.AutoMigrate(&Group{}) // 同步表
}
