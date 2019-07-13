package model

import (
	"fmt"
	"time"

	"github.com/fishjar/gin-rest-boilerplate/db"
	"github.com/fishjar/gin-rest-boilerplate/utils"
)

// User 用户模型
type User struct {
	Base
	UserName string     `json:"user_name" gorm:"type:VARCHAR(20);unique;not null" binding:"min=3,max=20"`
	UserType string     `json:"user_type" gorm:"type:VARCHAR(20);not null;default:'user'"`
	PassWord string     `json:"pass_word" gorm:"not null"`
	Name     *string    `json:"name" binding:"omitempty"`
	Age      *int       `json:"age" gorm:"type:TINYINT" binding:"omitempty,min=18"`
	Email    *string    `json:"email" binding:"omitempty,email"`
	Birth    *time.Time `json:"birth" gorm:"type:DATE" binding:"omitempty"`
}

// TableName 自定义用户表名
func (User) TableName() string {
	return "user"
}

func init() {
	db.DB.AutoMigrate(&User{})

	// 插入默认用户
	userName := "gabe"
	userType := "admin"
	passWord := "123456"
	user := User{
		UserName: userName,
		UserType: userType,
		PassWord: utils.MD5Pwd(passWord),
	}
	name := "Gabe Yuan"
	user.Name = &name
	// db.DB.Create(&user)
	if err := db.DB.Where(&user).FirstOrCreate(&user).Error; err != nil {
		fmt.Println("默认用户创建失败：", err)
	} else {
		fmt.Println("默认用户已创建")
		fmt.Println("用户名：", userName)
		fmt.Println("用户类型：", userType)
		fmt.Println("密码：", passWord)
	}
}
