/*
	模型定义
*/

package model

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Base 给所有模型共用
type Base struct {
	ID        uuid.UUID  `json:"id" gorm:"column:id;primary_key;not null"`                   // ID
	CreatedAt time.Time  `json:"createdAt" gorm:"column:created_at;not null"`                // 创建时间
	UpdatedAt time.Time  `json:"updatedAt" gorm:"column:update_at;not null"`                 // 更新时间
	DeletedAt *time.Time `json:"-" sql:"index" gorm:"column:deleted_at" binding:"omitempty"` // 软删除时间
}

// BeforeCreate 在创建前给ID赋值一个UUID
func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	return scope.SetColumn("ID", uuid)
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

// BulkDelete 批量删除
type BulkDelete struct {
	IDs []uuid.UUID `form:"ids" json:"ids" binding:"required"`
}

// BulkUpdate 批量更新
type BulkUpdate struct {
	IDs []uuid.UUID `form:"ids" json:"ids" binding:"required"`
	Obj interface{} `form:"obj" json:"obj" binding:"required"`
}
