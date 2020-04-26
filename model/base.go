/*
	模型定义
*/

package model

import (
	"fmt"
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
	uuid := uuid.NewV4()
	return scope.SetColumn("ID", uuid)
}

// PaginReq 分页查询参数
type PaginReq struct {
	Page uint   `form:"page,default=1"`
	Size uint   `form:"size,default=10"`
	Sort string `form:"sort,default=created_at desc"`
}

// PaginRes 分页查询结果
type PaginRes struct {
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
	IDs []uuid.UUID            `form:"ids" json:"ids" binding:"required"`
	Obj map[string]interface{} `form:"obj" json:"obj" binding:"required"`
}

// HTTPError 错误
type HTTPError struct {
	Code    int           `json:"code" binding:"required default:1"`        // 错误码
	Message string        `json:"message" binding:"required default:error"` // 错误提示
	Errors  []interface{} `json:"errors"`                                   // 详细错误信息
}

// HTTPSuccess http返回
type HTTPSuccess struct {
	Code    int         `json:"code" binding:"required default:0"`          // 状态码
	Message string      `json:"message" binding:"required default:success"` // 提示
	Data    interface{} `json:"data"`                                       // 返回数据
}

// BeforeUpdate 钩子
func (base *Base) BeforeUpdate() (err error) {
	fmt.Println("-------BeforeUpdate Hooks--------")
	fmt.Println(base.ID)
	return
}

// BeforeDelete 钩子
func (base *Base) BeforeDelete() (err error) {
	fmt.Println("-------BeforeDelete Hooks--------")
	fmt.Println(base.ID)
	return
}
