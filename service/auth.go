package service

import (
	"errors"

	"github.com/fishjar/gin-rest-boilerplate/db"
	"github.com/fishjar/gin-rest-boilerplate/model"
	"github.com/fishjar/gin-rest-boilerplate/utils"
)

// GetAuth 获取指定ID认证帐号
func GetAuth(id string) (model.Auth, error) {
	var auth model.Auth
	if err := db.DB.First(&auth, "id = ?", id).Error; err != nil {
		return auth, err
	}
	return auth, nil
}

// GetAuthWithRoles 获取指定ID认证帐号
func GetAuthWithRoles(id string) (model.Auth, error) {
	var auth model.Auth
	if err := db.DB.Preload("User").Preload("User.Roles").First(&auth, "id = ?", id).Error; err != nil {
		return auth, err
	}
	return auth, nil
}

// AuthCheck 检查Auth有效性
func AuthCheck(auth model.Auth) error {
	if !auth.IsEnabled {
		return errors.New("帐号已禁用") // 禁用
	}
	// TODO 过期时间检查
	// if (*auth.ExpireTime).Before(time.Now()) {
	// 	return errors.New("帐号已过期") // 过期
	// }
	return nil
}

// // AuthAndUserCheck 从数据库检查authID和userID有效性
// func AuthAndUserCheck(authID string, userID string) bool {
// 	if auth, err := GetAuth(authID); err != nil {
// 		return false // Auth不存在
// 	} else if err := AuthCheck(auth); err != nil {
// 		return false // 禁用或过期
// 	}
// 	if _, err := GetUser(userID); err != nil {
// 		return false // User不存在
// 	}
// 	return true
// }

// CreateAuthAccount 创建帐号
func CreateAuthAccount(data *model.AuthAccountCreateReq) error {
	// 开始事务
	tx := db.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback() // 回滚
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	// 创建用户
	user := model.User{
		Name:     data.UserName,
		Nickname: &data.Nickname,
		Mobile:   &data.Mobile,
	}
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 创建帐号
	passWord := utils.MD5Pwd(data.UserName, data.PassWord)
	auth := model.Auth{
		User:     &user,
		AuthType: "account",
		AuthName: data.UserName,
		AuthCode: &passWord,
	}
	if err := tx.Create(&auth).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
