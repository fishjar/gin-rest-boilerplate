package model

import (
	"time"

	"github.com/fishjar/gin-rest-boilerplate/db"
)

// Foo 定义模型
// gorm tags参考：https://gorm.io/docs/models.html
// binding tags参考：https://godoc.org/gopkg.in/go-playground/validator.v8
// 时间格式比较严格，参考：https://golang.org/pkg/time/#pkg-constants
// 模型定义中全部使用指针类型，是为了可以插入null值到数据库，但这样会造成一些使用的麻烦
// 也可以使用"database/sql"或"github.com/guregu/null"包中封装的类型
// 但是这样会造成binding验证失效，目前没有更好的实现办法，所以暂时全部使用指针类型
type Foo struct {
	Base
	Name     *string    `json:"name" gorm:"type:VARCHAR(20);unique;not null" binding:"min=3,max=20"` // 用户名
	Birthday *time.Time `json:"birthday" gorm:"type:DATE" binding:"omitempty"`                       // 生日
	GoodTime *time.Time `json:"good_time" gorm:"type:DATETIME;not null"`
	Age      *int       `json:"age" gorm:"type:TINYINT" binding:"omitempty,min=18,max=100"`
	Weight   *float32   `json:"weight" gorm:"type:FLOAT" binding:"omitempty,min=0.01,max=200"`
	Email    *string    `json:"email" gorm:"type:VARCHAR(255)" binding:"omitempty,email"`
	Homepage *string    `json:"homepage" gorm:"type:VARCHAR(255)" binding:"omitempty,url"`
	Notice   *string    `json:"notice" gorm:"type:TEXT" binding:"omitempty"`
	Intro    *string    `json:"intro" gorm:"type:TEXT" binding:"omitempty"`
	IsGood   *int       `json:"is_good" gorm:"type:TINYINT;default:1" binding:"omitempty,eq=0|eq=1"`
	MyExtra  *string    `json:"my_extra" gorm:"type:JSON" binding:"omitempty"`
	Status   *int       `json:"status" gorm:"type:TINYINT;not null;" binding:"eq=1|eq=2|eq=3"`
}

// TableName 自定义表名
func (Foo) TableName() string {
	return "foo"
}

func init() {
	db.DB.AutoMigrate(&Foo{}) // 同步表
}
