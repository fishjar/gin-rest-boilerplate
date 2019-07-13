package utils

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/fishjar/gin-rest-boilerplate/config"
)

// MD5Pwd 密码哈希函数
func MD5Pwd(str string) string {
	salt := config.PWDSalt
	m := md5.New()
	m.Write([]byte(str))
	m.Write([]byte(salt))
	return hex.EncodeToString(m.Sum(nil))
}
