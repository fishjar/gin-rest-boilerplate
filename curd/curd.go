/*
	curd 增删改查
*/

package curd

import (
	"fmt"
	"net/http"

	"github.com/fishjar/gin-rest-boilerplate/model"

	"github.com/fishjar/gin-rest-boilerplate/db"
	"github.com/fishjar/gin-rest-boilerplate/schema"
	"github.com/gin-gonic/gin"
)

// FindAndCountAll 分页查询
func FindAndCountAll(c *gin.Context, rows interface{}) {
	var q *schema.PaginQueryIn
	if err := c.ShouldBindQuery(&q); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
			"msg": "参数有误",
		})
		return
	}

	fmt.Println(q)

	offset := (q.Page - 1) * q.Size
	var total uint

	if err := db.DB.Model(&model.Auth{}).Count(&total).Limit(q.Size).Offset(offset).Order(q.Order).Preload("User").Find(&rows).Error; err != nil {
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
