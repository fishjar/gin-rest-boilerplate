/*
	服务封装
*/

package service

import (
	"net/http"

	"github.com/fishjar/gin-rest-boilerplate/model"
	"github.com/gin-gonic/gin"
)

// HTTPError 返回错误
func HTTPError(c *gin.Context, msg string, code int, err error) {
	c.JSON(code, model.HTTPError{
		Code:    code,
		Message: msg,
		Errors:  []error{err},
	})
}

// HTTPSuccess 返回成功
func HTTPSuccess(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, model.HTTPSuccess{
		Message: msg,
	})
}
