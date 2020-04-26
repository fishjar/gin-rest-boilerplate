package handler

import (
	"net/http"

	"github.com/fishjar/gin-rest-boilerplate/db"
	"github.com/fishjar/gin-rest-boilerplate/model"

	"github.com/gin-gonic/gin"
)

// RoleFindAndCountAll 查询多条信息
func RoleFindAndCountAll(c *gin.Context) {

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
	// var where model.Role
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
	var rows []model.Role

	// 查询数据
	if err := db.DB.Model(&rows).Where(where).Count(&total).Limit(q.Size).Offset(offset).Order(q.Sort).Preload("Users").Preload("Menus").Find(&rows).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err":     err.Error(),
			"message": "查询多条信息失败",
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, model.PaginRes{
		Page:  q.Page,
		Size:  q.Size,
		Total: total,
		Rows:  rows,
	})
}

// RoleFindByPk 根据主键查询单条信息
func RoleFindByPk(c *gin.Context) {

	// 获取参数
	id := c.Param("id")

	// 查询
	var data model.Role
	if err := db.DB.Preload("Users").Preload("Menus").First(&data, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err":     err.Error(),
			"message": "查询失败",
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, data)
}

// RoleSingleCreate 创建单条信息
func RoleSingleCreate(c *gin.Context) {

	// 绑定数据
	var data model.Role
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

// RoleUpdateByPk 更新单条信息
func RoleUpdateByPk(c *gin.Context) {

	// 获取参数
	id := c.Param("id")

	// 查询
	var data model.Role
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

// RoleDestroyByPk 删除单条信息
func RoleDestroyByPk(c *gin.Context) {

	// 获取参数
	id := c.Param("id")

	// 查询
	var data model.Role
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

// RoleFindOrCreate 查询或创建单条信息
func RoleFindOrCreate(c *gin.Context) {

	// 绑定数据
	var data model.Role
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

// RoleUpdateBulk 批量更新
func RoleUpdateBulk(c *gin.Context) {

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
	if err := db.DB.Model(model.Role{}).Where("id IN (?)", data.IDs).Updates(data.Obj).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err":     err.Error(),
			"message": "更新失败",
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, data)
}

// RoleDestroyBulk 批量删除
func RoleDestroyBulk(c *gin.Context) {

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
	if err := db.DB.Delete(model.Role{}, "id IN (?)", data.IDs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err":     err.Error(),
			"message": "删除失败",
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, data)
}
