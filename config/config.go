/*
	配置文件
*/

package config

import (
	"os"
	"strconv"
)

// 配置参数设置
// MySQL注意：
// 想要能正确的处理 time.Time，你需要添加 parseTime 参数。
// 想要完全的支持 UTF-8 编码，你需要修改charset=utf8 为 charset=utf8mb4。
// 如果你想指定主机，你需要使用 ()
const (
	MySQLURL     string = "root:123456@(localhost)/testdb?charset=utf8mb4&parseTime=True&loc=Local" // 数据库链接
	HTTPPort     int    = 4000                                                                      // 端口号
	JWTSignKey   string = "123456"                                                                  // JWT加密用的密钥
	JWTExpiresAt int    = 60 * 24                                                                   // JWT过期时间，分钟为单位
	PWDSalt      string = "123456"                                                                  // 密码哈希盐
)

// GetPort 获取端口号
func GetPort() int {
	port := HTTPPort
	if envPort := os.Getenv("PORT"); len(envPort) > 0 {
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
