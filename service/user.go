package service

import (
	"errors"

	"github.com/fishjar/gin-rest-boilerplate/db"
	"github.com/fishjar/gin-rest-boilerplate/model"
	"github.com/fishjar/gin-rest-boilerplate/utils"
	"github.com/gin-gonic/gin"
)

// GetUser 获取指定ID用户
func GetUser(id string) (model.User, error) {
	var user model.User
	if err := db.DB.First(&user, "id = ?", id).Error; err != nil {
		return user, err
	}
	return user, nil
}

// GetCurrentUser 获取当前用户
func GetCurrentUser(c *gin.Context) (model.User, error) {
	var user model.User

	// id := c.MustGet("UserID").(string) // 不存在的key会引发panic
	UserID, ok := c.Get("UserID")
	if !ok {
		return user, errors.New("没有登录")
	}

	id, ok := UserID.(string)
	if !ok {
		return user, errors.New("用户ID错误")
	}

	user, err := GetUser(id)
	if err != nil {
		return user, err
	}

	return user, nil
}

// GetUserRoles 获取用户角色列表
func GetUserRoles(user model.User) ([]model.Role, error) {
	var roles []model.Role

	if err := db.DB.Model(&user).Preload("Menus").Related(&roles, "Roles").Error; err != nil {
		return roles, err
	}

	return roles, nil
}

// GetUserMenus 获取用户菜单列表
func GetUserMenus(user model.User) ([]model.Menu, error) {
	var menus []model.Menu
	var tmpMenus []utils.IFUniqueItem

	roles, err := GetUserRoles(user)
	if err != nil {
		return menus, err
	}

	for _, role := range roles {
		for _, menu := range role.Menus {
			tmpMenus = append(tmpMenus, *menu)
		}
	}
	// menus = RemoveDuplicateMenu(menus) // 去重
	tmpMenus = utils.RemoveDuplicateElemt(tmpMenus)
	for _, v := range tmpMenus {
		menus = append(menus, v.(model.Menu))
	}

	return menus, nil
}

// func GetUserMenus(user model.User) ([]model.Menu, error) {
// 	var menus []model.Menu

// 	roles, err := GetUserRoles(user)
// 	if err != nil {
// 		return menus, err
// 	}

// 	for _, role := range roles {
// 		for _, menu := range role.Menus {
// 			menus = append(menus, *menu)
// 		}
// 	}
// 	menus = RemoveDuplicateMenu(menus) // 去重

// 	return menus, nil
// }

// GetCurrentUserRoles 获取当前用户角色列表
func GetCurrentUserRoles(c *gin.Context) ([]model.Role, error) {
	var roles []model.Role

	user, err := GetCurrentUser(c)
	if err != nil {
		return roles, err
	}

	roles, err = GetUserRoles(user)
	if err != nil {
		return roles, err
	}

	return roles, nil
}

// GetCurrentUserMenus 获取当前用户菜单列表
func GetCurrentUserMenus(c *gin.Context) ([]model.Menu, error) {
	var menus []model.Menu

	user, err := GetCurrentUser(c)
	if err != nil {
		return menus, err
	}

	menus, err = GetUserMenus(user)
	if err != nil {
		return menus, err
	}

	return menus, nil
}

// RemoveDuplicateMenu Menu去重
// 已用IFUniqueItem接口实现，此函数作废
func RemoveDuplicateMenu(menus []model.Menu) []model.Menu {
	result := make([]model.Menu, 0, len(menus))
	temp := map[string]struct{}{}
	for _, item := range menus {
		id := item.ID.String()
		if _, ok := temp[id]; !ok {
			temp[id] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
