package service

import (
	"errors"

	"github.com/fishjar/gin-rest-boilerplate/db"
	"github.com/fishjar/gin-rest-boilerplate/model"
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
