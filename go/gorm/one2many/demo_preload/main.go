package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 一对多关联查询 preload

// 第二步:定义一对多的表结构
// 2.1 定义一张 User 表
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"column:username"`
	// 添加外键关联
	CreditCards []CreditCard
}

// 2.2 定义一张 Card 表
type CreditCard struct {
	gorm.Model
	Number string
	UserID uint // 这个就是与 User 表关联的外键 名字是 结构体 + 主键(gorm的规定)
}

func main() {
	// 第一步: 创建 gorm 连接
	dsn := "root:iPanel@202204121853@tcp(106.55.94.54:33333)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(db)

	// 第三步: 通过 preload 方法进行一对多的查询
	// SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL
	users := []User{}
	db.Preload("CreditCards").Find(&users)
	strUser, _ := json.Marshal(&users)
	fmt.Println(string(strUser))
}
