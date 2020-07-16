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
	RedisPWD     string        = ""                                                                            // redis密码
	JWTSignKey   string        = "123456"                                                                      // JWT加密用的密钥
	JWTExpiresIn time.Duration = 60 * 24 * time.Minute                                                         // JWT过期时间
	PWDSalt      string        = "123456"                                                                      // 密码哈希盐
	SwagName     string        = "admin"                                                                       // swagger 用户名
	SwagPwd      string        = "123456"                                                                      // swagger 密码
)

// HTTPPort 端口号
var HTTPPort int = 4000

// GINENV 运行环境
var GINENV string = "dev"

// RedisAddr redis数据库链接
var RedisAddr string = "redis:6379"

// RootPath 项目根目录
var RootPath string

// UploadPath 项目根目录
var UploadPath string

func init() {
	if envPort := os.Getenv("GINPORT"); len(envPort) > 0 {
		portInt, err := strconv.Atoi(envPort)
		if err != nil {
			panic("获取端口失败")
		}
		HTTPPort = portInt
	}
}

func init() {
	if ginEnv := os.Getenv("GINENV"); len(ginEnv) > 0 {
		GINENV = ginEnv
	}
}

func init() {
	if GINENV == "dev" { // dev 环境使用本地redis
		RedisAddr = "localhost:6379"
	}
}

func init() {
	if rootPath, err := os.Getwd(); err != nil {
		RootPath = rootPath
	}
}

func init() {
	UploadPath := path.Join(RootPath, "tmp", "files")
	if err := os.MkdirAll(UploadPath, 0755); err != nil {
		panic("上传目录创建失败")
	}
}
