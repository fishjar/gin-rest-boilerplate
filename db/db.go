/*
	db 数据库连接
*/

package db

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/fishjar/gin-rest-boilerplate/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // 引入sqlite驱动
	// _ "github.com/jinzhu/gorm/dialects/mysql" // 引入mysql驱动
)

// DB 为ORM全局实例
var DB *gorm.DB

func init() {
	// 默认MYSQL
	dbDriver := "mysql"
	dbPath := config.MySQLURL

	if env := config.GetEnv(); env == "dev" {
		// dev环境使用sqlite
		dbDriver = "sqlite3"
		rootPath, _ := os.Getwd()
		dbDir := path.Join(rootPath, "tmp/db")
		// 创建数据库目录
		err := os.MkdirAll(dbDir, 0755)
		if err != nil {
			panic("创建数据库目录失败")
		}
		dbPath = path.Join(dbDir, "sqlite.db")
	}

	db, err := gorm.Open(dbDriver, dbPath)
	if err != nil {
		fmt.Println("打开数据库错误：", err)
		panic("连接数据库失败")
	}

	db.LogMode(true) // 生产环境建议关闭

	db.DB()
	db.DB().Ping()

	// 链接池设置
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Hour)

	DB = db
}
