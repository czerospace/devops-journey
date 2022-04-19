package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// 自定义第三张表

// 第二步: 自定义第三张表
/*
	模拟美团外卖地址
	一个用户可以填写多个地址
	一个地址可能被多个用户填写
*/

type Person struct {
	ID   int
	Name string `gorm:"many2many:person_addresses"`
}
type Address struct {
	ID   uint
	Name string
}

// 第三张表自定义

type PersonAddress struct {
	PersonID  int `gorm:"primaryKey"` // 对应表的主键
	AddressId int `gorm:"primaryKey"`
	CreateAt  time.Time
}

func main() {
	// 第一步: 连接数据库
	dsn := "root:iPanel@202204121853@tcp(106.55.94.54:33333)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println(db)

	// 第三步: 创建第三张表
	db.AutoMigrate(Person{}, Address{}, PersonAddress{})
}
