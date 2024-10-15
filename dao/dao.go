package dao

import (
	"fmt"
	"go-gin-first/config"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Db  *gorm.DB
	err error
)

func Init() {
	// 尝试连接数据库
	Db, err = gorm.Open(mysql.Open(config.MysqlDb), &gorm.Config{})
	if err != nil {
		log.Fatalf("数据库连接错误！%v", err) // 使用 log.Fatal 记录错误并退出
		return
	}

	// 获取底层数据库连接
	sqlDB, err := Db.DB()
	if err != nil {
		log.Fatalf("获取数据库连接错误！%v", err) // 使用 log.Fatal 记录错误并退出
		return
	}

	// Ping 数据库以检查连接
	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("数据库无法 ping 通！%v", err) // 使用 log.Fatal 记录错误并退出
		return
	}

	// 连接成功，输出提示
	fmt.Println("数据库连接成功")
}
