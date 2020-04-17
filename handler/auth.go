package handler

import (
	"net/http"

	"github.com/fishjar/gin-rest-boilerplate/db"
	"github.com/fishjar/gin-rest-boilerplate/model"
	"github.com/fishjar/gin-rest-boilerplate/schema"

	"github.com/gin-gonic/gin"
)

// AuthFindAndCountAll 查询多条信息
func AuthFindAndCountAll(c *gin.Context) {

	// 获取参数
	// ShouldBindQuery
	// db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)
	// pageNum, _ := strconv.Atoi(c.DefaultQuery("page_num", "1"))    // 页码
	// pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10")) // 每页数目
	// order := c.DefaultQuery("sorter", "")                          // 排序
	// where := c.DefaultQuery("where", "")                           // 检索条件 c.QueryMap("querys")
	// offset := (pageNum - 1) * pageSize

	// // 查询数据
	// var rows []model.Auth
	// var count uint
	// if err := db.DB.Model(&rows).Where(where).Count(&count).Limit(pageSize).Offset(offset).Order(order).Preload("User").Find(&rows).Error; err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"err":     err,
	// 		"message": "查询多条信息失败",
	// 	})
	// 	return
	// }

	// // 返回数据
	// c.JSON(http.StatusOK, gin.H{
	// 	"rows":  rows,
	// 	"count": count,
	// })
	// var rows []model.Auth
	// curd.FindAndCountAll(c, model.Auth)

	var q *schema.PaginQueryIn
	if err := c.ShouldBindQuery(&q); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
			"msg": "参数有误",
		})
		return
	}

	params := c.QueryMap("params")
	offset := (q.Page - 1) * q.Size
	var total uint
	var rows []model.Auth
	where := make(map[string]interface{})
	for k, v := range params {
		where[k] = v
	}

	if err := db.DB.Model(&rows).Where(params).Count(&total).Limit(q.Size).Offset(offset).Order(q.Order).Preload("User").Find(&rows).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err,
			"msg": "查询多条信息失败",
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, schema.PaginQueryOut{
		Page:  q.Page,
		Size:  q.Size,
		Total: total,
		Rows:  rows,
	})
}

// AuthFindByPk 根据主键查询单条信息
func AuthFindByPk(c *gin.Context) {

	// 获取参数
	// c.ShouldBindUri
	id := c.Param("id")

	// 查询
	var data model.Auth
	if err := db.DB.Where("id = ?", id).Preload("User").First(&data).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err":     err,
			"message": "根据主键查询单条信息失败",
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, data)
}

// AuthSingleCreate 创建单条信息
func AuthSingleCreate(c *gin.Context) {

	// 绑定数据
	var data model.Auth
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     err,
			"message": "数据绑定失败",
		})
		return
	}

	// 插入数据
	if err := db.DB.Create(&data).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err":     err,
			"message": "插入数据失败",
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, data)
}

// AuthUpdateByPk 更新单条信息
func AuthUpdateByPk(c *gin.Context) {

	// 获取参数
	id := c.Param("id")

	// 查询
	var data model.Auth
	if err := db.DB.Where("id = ?", id).First(&data).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err":     err,
			"message": "查询失败",
		})
		return
	}

	// 绑定新数据
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     err,
			"message": "数据绑定失败",
		})
		return
	}

	// 更新数据
	if err := db.DB.Model(&data).Updates(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     err,
			"message": "更新失败",
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, data)
}

// AuthDestroyByPk 删除单条信息
func AuthDestroyByPk(c *gin.Context) {

	// 获取参数
	id := c.Param("id")

	// 查询
	var data model.Auth
	if err := db.DB.Where("id = ?", id).First(&data).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err":     err,
			"message": "查询失败",
		})
		return
	}

	// 删除
	if err := db.DB.Delete(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     err,
			"message": "删除失败",
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, data)
}

// AuthFindOrCreate 查询或创建单条信息
func AuthFindOrCreate(c *gin.Context) {

	// 绑定数据
	var data model.Auth
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     err,
			"message": "数据绑定失败",
		})
		return
	}

	// 插入数据
	if err := db.DB.Where(&data).FirstOrCreate(&data).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err":     err,
			"message": "查询或创建数据失败",
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, data)
}
