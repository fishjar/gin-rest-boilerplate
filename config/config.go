/*
	配置文件
*/

package config

import (
	"os"
	"strconv"
)

// 配置参数设置
const (
	MySQLURL     string = "root:123456@/testdb" // 数据库链接
	HTTPPort     int    = 4000                  // 端口号
	JWTSignKey   string = "123456"              // JWT加密用的密钥
	JWTExpiresAt int    = 60 * 24               // JWT过期时间，分钟为单位
	PWDSalt      string = "123456"              // 密码哈希盐
)

// GetPort 获取端口号
func GetPort() int {
	port := HTTPPort
	envPort := os.Getenv("PORT")
	if len(envPort) > 0 {
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
	ginEnv := os.Getenv("GINENV")
	if len(ginEnv) > 0 {
		env = ginEnv
	}
	return env
}
