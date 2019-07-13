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
	ID        uuid.UUID  `json:"id" gorm:"primary_key;"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"update_at"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
}

// BeforeCreate 在创建前给ID赋值一个UUID
func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	return scope.SetColumn("ID", uuid)
}
