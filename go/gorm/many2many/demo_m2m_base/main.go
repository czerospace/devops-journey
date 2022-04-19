package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 第二步: 定义多对多的表结构

type User struct {
	gorm.Model
	// 建立和 Language 的关联关系
	/*
		Languages 字段名: 自己定义的，可以随意
		gorm:"many2many 指定了 User 表和 Language 表是多对多的关系
		user_languages 第三张表的表名,存放的是多对多的关系(自动生成第三张表)
	*/
	Languages []Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	gorm.Model
	Name string
}

func main() {
	// 第一步: 连接数据库
	dsn := "root:iPanel@202204121853@tcp(106.55.94.54:33333)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println(db)

	// 第三步： 创建多对多表结构
	db.AutoMigrate(User{}, Language{})
}
