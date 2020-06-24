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
	_ "github.com/jinzhu/gorm/dialects/mysql"  // 引入mysql驱动
	_ "github.com/jinzhu/gorm/dialects/sqlite" // 引入sqlite驱动
)

// DB 为ORM全局实例
var DB *gorm.DB

func init() {
	// 获取数据库参数
	dbDriver := "mysql"                       // 数据库驱动，默认MYSQL
	dbPath := config.MySQLAddr                // 数据库地址
	if env := config.GetEnv(); env == "dev" { // dev环境使用sqlite
		dbDriver = "sqlite3"
		rootPath, _ := os.Getwd()
		dbDir := path.Join(rootPath, "tmp", "db")
		if err := os.MkdirAll(dbDir, 0755); err != nil {
			fmt.Println(err.Error())
			panic("创建数据库目录失败")
		}
		dbPath = path.Join(dbDir, "sqlite.db")
	}

	// 创建数据库连接
	db, err := gorm.Open(dbDriver, dbPath)
	if err != nil {
		fmt.Println("打开数据库错误：", err.Error())
		panic("连接数据库失败")
	}
	db.LogMode(true) // 生产环境建议关闭

	// 测试数据库连接
	if db.DB() == nil { // 如果数据库底层连接的不是一个 *sql.DB，那么该方法会返回 nil
		fmt.Println("获取数据库接口错误")
		panic("连接数据库失败")
	}
	if err := db.DB().Ping(); err != nil {
		fmt.Println("Ping数据库错误：", err.Error())
		panic("连接数据库失败")
	}

	// db设置
	db.DB().SetMaxIdleConns(10)           // 设置连接池中的最大闲置连接数
	db.DB().SetMaxOpenConns(100)          // 设置数据库的最大连接数量
	db.DB().SetConnMaxLifetime(time.Hour) // 设置连接的最大可复用时间
	// db.SetLogger(log.New(os.Stdout, "\r\n", 0)) // log设置

	DB = db
}
