package handler

import (
	"net/http"
	"strconv"

	"github.com/fishjar/gin-rest-boilerplate/db"
	"github.com/fishjar/gin-rest-boilerplate/model"

	"github.com/gin-gonic/gin"
)

// UserGroupFindAndCountAll 查询多条信息
func UserGroupFindAndCountAll(c *gin.Context) {

	// 获取参数
	pageNum, _ := strconv.Atoi(c.DefaultQuery("page_num", "1"))    // 页码
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10")) // 每页数目
	order := c.DefaultQuery("sorter", "")                          // 排序
	where := c.DefaultQuery("where", "")                           // 检索条件
	offset := (pageNum - 1) * pageSize

	// 查询数据
	var rows []model.UserGroup
	var count uint
	if err := db.DB.Model(&rows).Where(where).Count(&count).Limit(pageSize).Offset(offset).Order(order).Find(&rows).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err":     err,
			"message": "查询失败",
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"rows":  rows,
		"count": count,
	})
}

// UserGroupFindByPk 根据主键查询单条信息
func UserGroupFindByPk(c *gin.Context) {

	// 获取参数
	id := c.Param("id")

	// 查询
	var data model.UserGroup
	if err := db.DB.Where("id = ?", id).First(&data).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err":     err,
			"message": "查询失败",
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, data)
}

// UserGroupSingleCreate 创建单条信息
func UserGroupSingleCreate(c *gin.Context) {

	// 绑定数据
	var data model.UserGroup
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

// UserGroupUpdateByPk 更新单条信息
func UserGroupUpdateByPk(c *gin.Context) {

	// 获取参数
	id := c.Param("id")

	// 查询
	var data model.UserGroup
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

// UserGroupDestroyByPk 删除单条信息
func UserGroupDestroyByPk(c *gin.Context) {

	// 获取参数
	id := c.Param("id")

	// 查询
	var data model.UserGroup
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

// UserGroupFindOrCreate 查询或创建单条信息
func UserGroupFindOrCreate(c *gin.Context) {

	// 绑定数据
	var data model.UserGroup
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
