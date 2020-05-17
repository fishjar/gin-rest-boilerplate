package service

import (
	"errors"
	"mime/multipart"
	"net/http"
	"os"
	"path"

	"github.com/fishjar/gin-rest-boilerplate/config"
	"github.com/fishjar/gin-rest-boilerplate/model"
	"github.com/fishjar/gin-rest-boilerplate/utils"
	"github.com/gin-gonic/gin"
)

// SaveFile 保存文件
func SaveFile(c *gin.Context, file *multipart.FileHeader) (model.UploadRes, error) {
	// 判断文件类型
	contentType := GetFileHeaderContentType(file)

	// 获取文件扩展名
	srcName := file.Filename
	extName := path.Ext(srcName)
	if len(extName) == 0 { // 限制没有文件扩展名的文件
		return model.UploadRes{}, errors.New("文件扩展名有误")
	}

	// 计算文件md5
	md5str, err := utils.MD5File(file)
	if err != nil {
		return model.UploadRes{}, err
	}

	// 判断文件是否已经存在
	dstPath := path.Join(config.GetFileDir(), md5str[0:2], md5str[2:4])
	fileName := md5str + extName
	filePath := path.Join(md5str[0:2], md5str[2:4], fileName)
	dst := path.Join(dstPath, fileName)
	if FileExist(dst) {
		return model.UploadRes{
			File:  srcName,
			Ext:   extName,
			Name:  fileName,
			Path:  filePath,
			Type:  contentType,
			Isnew: false,
		}, nil
	}

	// 创建目录
	if err := os.MkdirAll(dstPath, 0755); err != nil {
		return model.UploadRes{}, err
	}

	// 保存文件
	if err := c.SaveUploadedFile(file, dst); err != nil {
		return model.UploadRes{}, err
	}

	// TODO: 如果是图片格式，压缩及剪裁

	return model.UploadRes{
		File:  srcName,
		Ext:   extName,
		Name:  fileName,
		Path:  filePath,
		Type:  contentType,
		Isnew: true,
	}, nil
}

// GetFileContentType 获取文件类型
func GetFileContentType(file *os.File) (string, error) {

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := file.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}

// GetFileHeaderContentType 获取文件类型
func GetFileHeaderContentType(file *multipart.FileHeader) string {
	buffer := make([]byte, 512)
	tmpFile, _ := file.Open()
	defer tmpFile.Close()
	tmpFile.Read(buffer)
	contentType := http.DetectContentType(buffer)
	return contentType
}

// FileExist 检查文件或目录是否存在
func FileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
