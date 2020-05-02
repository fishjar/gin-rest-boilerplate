package service

import (
	"errors"
	"mime/multipart"
	"os"
	"path"

	"github.com/fishjar/gin-rest-boilerplate/config"
	"github.com/fishjar/gin-rest-boilerplate/model"
	"github.com/fishjar/gin-rest-boilerplate/utils"
	"github.com/gin-gonic/gin"
)

// SaveFile 保存文件
func SaveFile(c *gin.Context, file *multipart.FileHeader) (model.UploadRes, error) {
	// 获取文件扩展名
	srcName := file.Filename
	extName := path.Ext(srcName)
	if len(extName) == 0 {
		return model.UploadRes{
			File: srcName,
			Ext:  extName,
		}, errors.New("文件格式有误")
	}

	// 计算文件md5
	md5str, err := utils.MD5File(file)
	if err != nil {
		return model.UploadRes{
			File: srcName,
			Ext:  extName,
		}, err
	}

	// 创建目录
	dstPath := path.Join(config.GetFileDir(), md5str[0:2], md5str[2:4])
	if err := os.MkdirAll(dstPath, 0755); err != nil {
		return model.UploadRes{
			File: srcName,
			Ext:  extName,
		}, err
	}

	// 保存文件
	fileName := md5str + extName
	filePath := path.Join(md5str[0:2], md5str[2:4], fileName)
	dst := path.Join(dstPath, fileName)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		return model.UploadRes{
			File: srcName,
			Ext:  extName,
			Name: fileName,
		}, err
	}

	return model.UploadRes{
		File: srcName,
		Ext:  extName,
		Name: fileName,
		Path: filePath,
	}, nil
}
