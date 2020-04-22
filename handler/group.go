package handler

import (
	"net/http"

	"github.com/fishjar/gin-rest-boilerplate/db"
	"github.com/fishjar/gin-rest-boilerplate/model"

	"github.com/gin-gonic/gin"
)

// GroupFindAndCountAll 查询多条信息
func GroupFindAndCountAll(c *gin.Context) {

	// 参数绑定
	var q *model.PaginQueryIn
	if err := c.ShouldBindQuery(&q); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
			"msg": "参数有误",
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
	// var where model.Group
	// if err := mapstructure.Decode(params, &where); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"err": err.Error(),
	// 		"msg": "查询参数有误",
	// 	})
	// 	return
	// }

	// 分页参数
	offset := (q.Page - 1) * q.Size
	var total uint
	var rows []model.Group

	// 查询数据
	if err := db.DB.Model(&rows).Where(where).Count(&total).Limit(q.Size).Offset(offset).Order(q.Sort).Preload("Leader").Preload("Users").Find(&rows).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
			"msg": "查询多条信息失败",
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, model.PaginQueryOut{
		Page:  q.Page,
		Size:  q.Size,
		Total: total,
		Rows:  rows,
	})
}

// GroupFindByPk 根据主键查询单条信息
func GroupFindByPk(c *gin.Context) {

	// 获取参数
	id := c.Param("id")

	// 查询
	var data model.Group
	if err := db.DB.Preload("Leader").Preload("Users").First(&data, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err": err.Error(),
			"msg": "查询失败",
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, data)
}

// GroupSingleCreate 创建单条信息
func GroupSingleCreate(c *gin.Context) {

	// 绑定数据
	var data model.Group
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
			"msg": "数据绑定失败",
		})
		return
	}

	// 插入数据
	if err := db.DB.Create(&data).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
			"msg": "插入数据失败",
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, data)
}

// GroupUpdateByPk 更新单条信息
func GroupUpdateByPk(c *gin.Context) {

	// 获取参数
	id := c.Param("id")

	// 查询
	var data model.Group
	if err := db.DB.First(&data, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err": err.Error(),
			"msg": "查询失败",
		})
		return
	}

	// 绑定新数据
	var obj map[string]interface{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
			"msg": "数据绑定失败",
		})
		return
	}

	// 更新数据
	if err := db.DB.Model(&data).Updates(obj).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
			"msg": "更新失败",
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, data)
}

// GroupDestroyByPk 删除单条信息
func GroupDestroyByPk(c *gin.Context) {

	// 获取参数
	id := c.Param("id")

	// 查询
	var data model.Group
	if err := db.DB.Where("id = ?", id).First(&data).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err": err.Error(),
			"msg": "查询失败",
		})
		return
	}

	// 删除
	if err := db.DB.Delete(&data).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
			"msg": "删除失败",
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, data)
}

// GroupFindOrCreate 查询或创建单条信息
func GroupFindOrCreate(c *gin.Context) {

	// 绑定数据
	var data model.Group
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
			"msg": "数据绑定失败",
		})
		return
	}

	// 插入数据
	if err := db.DB.FirstOrCreate(&data, data).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
			"msg": "查询或创建数据失败",
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, data)
}

// GroupUpdateBulk 批量更新
func GroupUpdateBulk(c *gin.Context) {

	var data model.BulkUpdate

	// 绑定数据
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
			"msg": "数据绑定失败",
		})
		return
	}

	// 判断ID列表是否为空
	// if len(data.IDs) == 0 {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"msg": "ids列表不能空",
	// 	})
	// 	return
	// }

	// 更新数据
	if err := db.DB.Model(model.Group{}).Where("id IN (?)", data.IDs).Updates(data.Obj).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
			"msg": "更新失败",
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, data)
}

// GroupDestroyBulk 批量删除
func GroupDestroyBulk(c *gin.Context) {

	var data model.BulkDelete

	// 绑定数据
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
			"msg": "数据绑定失败",
		})
		return
	}

	// 删除数据
	if err := db.DB.Delete(model.Group{}, "id IN (?)", data.IDs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
			"msg": "删除失败",
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, data)
}
