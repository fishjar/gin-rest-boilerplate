package handler

import (
	"net/http"

	"github.com/fishjar/gin-rest-boilerplate/model"
	"github.com/fishjar/gin-rest-boilerplate/service"
	"github.com/gin-gonic/gin"
)

// UploudFile 文件上传，单文件
// @Summary				文件上传，单文件
// @Description			文件上传，单文件...
// @Tags				admin
// @Accept				mpfd
// @Produce				json
// @Param				file formData file true "文件"
// @Success				200 {object} model.UploadSuccess
// @Failure 			500 {object} model.HTTPError
// @Router				/admin/upload/file [post]
// @Security			ApiKeyAuth
func UploudFile(c *gin.Context) {
	// 获取文件
	file, err := c.FormFile("file")
	if err != nil {
		service.HTTPError(c, "文件上传失败：没有获取到文件", http.StatusBadRequest, err)
		return
	}

	data, err := service.SaveFile(c, file)
	if err != nil {
		service.HTTPError(c, "文件上传失败：保存文件失败", http.StatusInternalServerError, err)
		return
	}

	// 返回成功
	c.JSON(http.StatusOK, model.UploadSuccess{
		HTTPSuccess: model.HTTPSuccess{
			Message: "上传成功",
		},
		Data: data,
	})
}
