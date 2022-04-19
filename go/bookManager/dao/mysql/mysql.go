package mysql

import (
	"bookManager/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB 为什么要将变量定义在全局,是因为在其他包里需要使用
var DB *gorm.DB

// InitMysql 初始化 mysql 链接
func InitMysql() {
	// 1.链接数据库
	dsn := "root:iPanel@202204121853@tcp(106.55.94.54:33333)/books?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db

	err = DB.AutoMigrate(model.User{}, model.Book{})
	if err != nil {
		fmt.Println("创建表结构失败", err)
	}
}
