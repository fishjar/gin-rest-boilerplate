package handler

import (
	"net/http"

	"github.com/fishjar/gin-rest-boilerplate/db"
	"github.com/fishjar/gin-rest-boilerplate/model"
	"github.com/fishjar/gin-rest-boilerplate/service"

	"github.com/gin-gonic/gin"
)

// UserFindAndCountAll 查询多条信息
// @Summary				用户列表
// @Description			用户列表...
// @Tags				user
// @Accept				json
// @Produce				json
// @Param				q query model.PaginReq false "name search by q"
// @Success				200 {object} model.UserListRes
// @Failure 			400 {object} httputil.HTTPError
// @Failure 			404 {object} httputil.HTTPError
// @Router				/admin/users [get]
// @Security			ApiKeyAuth
func UserFindAndCountAll(c *gin.Context) {

	// 参数绑定
	var q *model.PaginReq
	if err := c.ShouldBindQuery(&q); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     err.Error(),
			"message": "参数有误",
		})
		return
	}

	// 条件参数
	params := c.QueryMap("params")
	// map 参数，map[string]T 必须转为 map[string]interface{}
	where := make(map[string]interface{}, len(params))
	for k, v := range params {
		// 有模型不存在的key时，后面的查询会报错，需要过滤掉
		if k == "name" || k == "gender" {
			where[k] = v
		}
	}
	// struct 参数，后面使用指针
	// var where model.User
	// if err := mapstructure.Decode(params, &where); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"err": err.Error(),
	// 		"message": "查询参数有误",
	// 	})
	// 	return
	// }

	// 分页参数
	offset := (q.Page - 1) * q.Size
	var total uint
	var rows []model.User

	// 查询数据
	if err := db.DB.Model(&rows).Where(where).Count(&total).Limit(q.Size).Offset(offset).Order(q.Sort).Preload("Auths").Preload("Roles").Preload("Groups").Preload("Friends").Find(&rows).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err":     err.Error(),
			"message": "查询多条信息失败",
		})
		return
	}

	// 返回数据
	// c.JSON(http.StatusOK, model.PaginRes{
	// 	Page:  q.Page,
	// 	Size:  q.Size,
	// 	Total: total,
	// 	Rows:  rows,
	// })
	c.JSON(http.StatusOK, model.UserListRes{
		PaginRes: model.PaginRes{
			Page:  q.Page,
			Size:  q.Size,
			Total: total,
		},
		Rows: rows,
	})
}

// UserFindByPk 根据主键查询单条信息
func UserFindByPk(c *gin.Context) {

	// 获取参数
	id := c.Param("id")

	// 查询
	var data model.User
	if err := db.DB.Preload("Auths").Preload("Roles").Preload("Groups").Preload("Friends").First(&data, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err":     err.Error(),
			"message": "查询失败",
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, data)
}

// UserSingleCreate 创建单条信息
func UserSingleCreate(c *gin.Context) {

	// 绑定数据
	var data model.User
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     err.Error(),
			"message": "数据绑定失败",
		})
		return
	}

	// 插入数据
	if err := db.DB.Create(&data).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err":     err.Error(),
			"message": "插入数据失败",
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, data)
}

// UserUpdateByPk 更新单条信息
func UserUpdateByPk(c *gin.Context) {

	// 获取参数
	id := c.Param("id")

	// 查询
	var data model.User
	if err := db.DB.First(&data, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err":     err.Error(),
			"message": "查询失败",
		})
		return
	}

	// 绑定新数据
	var obj map[string]interface{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     err.Error(),
			"message": "数据绑定失败",
		})
		return
	}

	// 更新数据
	if err := db.DB.Model(&data).Updates(obj).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err":     err.Error(),
			"message": "更新失败",
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, data)
}

// UserDestroyByPk 删除单条信息
func UserDestroyByPk(c *gin.Context) {

	// 获取参数
	id := c.Param("id")

	// 查询
	var data model.User
	if err := db.DB.Where("id = ?", id).First(&data).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err":     err.Error(),
			"message": "查询失败",
		})
		return
	}

	// 删除
	if err := db.DB.Delete(&data).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err":     err.Error(),
			"message": "删除失败",
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, data)
}

// UserFindOrCreate 查询或创建单条信息
func UserFindOrCreate(c *gin.Context) {

	// 绑定数据
	var data model.User
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     err.Error(),
			"message": "数据绑定失败",
		})
		return
	}

	// 插入数据
	if err := db.DB.FirstOrCreate(&data, data).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err":     err.Error(),
			"message": "查询或创建数据失败",
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, data)
}

// UserUpdateBulk 批量更新
func UserUpdateBulk(c *gin.Context) {

	var data model.BulkUpdate

	// 绑定数据
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     err.Error(),
			"message": "数据绑定失败",
		})
		return
	}

	// 判断ID列表是否为空
	// if len(data.IDs) == 0 {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"message": "ids列表不能空",
	// 	})
	// 	return
	// }

	// 更新数据
	if err := db.DB.Model(model.User{}).Where("id IN (?)", data.IDs).Updates(data.Obj).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err":     err.Error(),
			"message": "更新失败",
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, data)
}

// UserDestroyBulk 批量删除
func UserDestroyBulk(c *gin.Context) {

	var data model.BulkDelete

	// 绑定数据
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     err.Error(),
			"message": "数据绑定失败",
		})
		return
	}

	// 删除数据
	if err := db.DB.Delete(model.User{}, "id IN (?)", data.IDs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err":     err.Error(),
			"message": "删除失败",
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, data)
}

// UserFindMyRoles 查找本人角色
func UserFindMyRoles(c *gin.Context) {
	// 获取当前用户角色列表
	roles, err := service.GetCurrentUserRoles(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err":     err.Error(),
			"message": "查询角色失败",
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, roles)
}

// UserFindMyMenus 查找本人菜单
func UserFindMyMenus(c *gin.Context) {
	// 获取当前用户菜单列表
	menus, err := service.GetCurrentUserMenus(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err":     err.Error(),
			"message": "查询菜单失败",
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, menus)
}
