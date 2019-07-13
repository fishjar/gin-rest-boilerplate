package handler

import (
	"net/http"
	"strconv"

	"github.com/fishjar/gin-rest-boilerplate/db"
	"github.com/fishjar/gin-rest-boilerplate/model"

	"github.com/gin-gonic/gin"
)

// FindFoos 查询多条信息
func FindFoos(c *gin.Context) {

	// // 获取全局用户信息示例
	// if userName, ok := c.Get("userName"); ok {
	// 	fmt.Println("userName", userName)
	// }
	// if userType, ok := c.Get("userType"); ok {
	// 	fmt.Println("userType", userType)
	// }

	// 获取参数
	pageNum, _ := strconv.Atoi(c.DefaultQuery("page_num", "1"))    // 页码
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10")) // 每页数目
	order := c.DefaultQuery("sorter", "")                          // 排序
	where := c.DefaultQuery("where", "")                           // 检索条件
	offset := (pageNum - 1) * pageSize

	// 查询数据
	var foos []model.Foo
	var count uint
	if err := db.DB.Model(&model.Foo{}).Count(&count).Limit(pageSize).Offset(offset).Where(where).Order(order).Find(&foos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err":     err,
			"message": "查询失败",
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"rows":  foos,
		"count": count,
	})
}

// FindFooByID 根据ID查询单条信息
func FindFooByID(c *gin.Context) {

	// 获取参数
	id := c.Param("id")

	// 查询
	var foo model.Foo
	if err := db.DB.Where("id = ?", id).First(&foo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err":     err,
			"message": "查询失败",
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, foo)
}

// FindOrCreateFoo 查询或创建单条信息
func FindOrCreateFoo(c *gin.Context) {

	// 绑定数据
	var foo model.Foo
	if err := c.ShouldBind(&foo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     err,
			"message": "数据绑定失败",
		})
		return
	}

	// 插入数据
	if err := db.DB.Where(&foo).FirstOrCreate(&foo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err":     err,
			"message": "查询或创建数据失败",
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, foo)
}

// CreateFoo 创建单条信息
func CreateFoo(c *gin.Context) {

	// 绑定数据
	var foo model.Foo
	if err := c.ShouldBind(&foo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     err,
			"message": "数据绑定失败",
		})
		return
	}

	// 插入数据
	if err := db.DB.Create(&foo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err":     err,
			"message": "插入数据失败",
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, foo)
}

// UpdateFoo 更新单条信息
func UpdateFoo(c *gin.Context) {

	// 获取参数
	id := c.Param("id")

	// 查询
	var foo model.Foo
	if err := db.DB.Where("id = ?", id).First(&foo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err":     err,
			"message": "查询失败",
		})
		return
	}

	// 绑定新数据
	var fooNew model.Foo
	if err := c.ShouldBind(&fooNew); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     err,
			"message": "数据绑定失败",
		})
		return
	}

	// 更新数据
	if err := db.DB.Model(&foo).Updates(&fooNew).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     err,
			"message": "更新失败",
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, foo)
}

// DeleteFoo 删除单条信息
func DeleteFoo(c *gin.Context) {

	// 获取参数
	id := c.Param("id")

	// 查询
	var foo model.Foo
	if err := db.DB.Where("id = ?", id).First(&foo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err":     err,
			"message": "查询失败",
		})
		return
	}

	// 删除
	if err := db.DB.Delete(&foo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     err,
			"message": "删除失败",
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, foo)
}
