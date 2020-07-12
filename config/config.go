/*
	配置文件
*/

package config

import (
	"os"
	"path"
	"strconv"
	"time"
)

// 配置参数设置
// MySQL注意：
// 想要能正确的处理 time.Time，你需要添加 parseTime 参数。
// 想要完全的支持 UTF-8 编码，你需要修改charset=utf8 为 charset=utf8mb4。
// 如果你想指定主机，你需要使用 ()
const (
	MySQLAddr    string        = "root:123456@tcp(mysql:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local" // 数据库链接
	RedisAddr    string        = "redis:6379"                                                                  // redis数据库链接
	RedisPWD     string        = ""                                                                            // redis密码
	HTTPPort     int           = 4000                                                                          // 端口号
	JWTSignKey   string        = "123456"                                                                      // JWT加密用的密钥
	JWTExpiresIn time.Duration = 60 * 24 * time.Minute                                                         // JWT过期时间
	PWDSalt      string        = "123456"                                                                      // 密码哈希盐
	SwagName     string        = "admin"                                                                       // swagger 用户名
	SwagPwd      string        = "123456"                                                                      // swagger 密码
)

// GetPort 获取端口号
func GetPort() int {
	port := HTTPPort
	if envPort := os.Getenv("GINPORT"); len(envPort) > 0 {
		portInt, err := strconv.Atoi(envPort)
		if err != nil {
			panic("获取端口失败")
		}
		port = portInt
	}
	return port
}

// GetEnv 获取运行环境
func GetEnv() string {
	env := "dev"
	if ginEnv := os.Getenv("GINENV"); len(ginEnv) > 0 {
		env = ginEnv
	}
	return env
}

// GetRedisAddr 获取Redis地址
func GetRedisAddr() string {
	if env := GetEnv(); env == "dev" { // dev 环境使用本地redis
		return "localhost:6379"
	}
	return RedisAddr
}

// GetFileDir 获取文件上传目录
func GetFileDir() string {
	rootPath, _ := os.Getwd()
	fileDir := path.Join(rootPath, "tmp", "files")
	return fileDir
}

// // 返回的状态码
// const (
// 	CODEOK = 0 // 一切正常
// )
