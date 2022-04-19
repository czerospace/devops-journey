package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 一对多

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

	// 第三步: 创建表结构
	db.AutoMigrate(User{}, CreditCard{})

	// 第四步: 创建一对多
	// INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`username`) VALUES ('2022-04-13 17:08:02.815','2022-04-13 17:08:02.815',NULL,'miaozong')
	user := User{
		Username: "miaozong",
		CreditCards: []CreditCard{
			{
				Number: "0001",
			},
			{
				Number: "0002",
			},
		},
	}
	db.Create(&user)

	// 第五步: 给 miaozong 添加信用卡
	// 5.1 查找 miaozong
	u := User{Username: "miaozong"}
	db.First(&u)
	// 5.2 对 miaozong 进行操作
	// UPDATE `users` SET `updated_at`='2022-04-13 17:08:03.48' WHERE `users`.`deleted_at` IS NULL AND `id` = 1
	db.Model(&u).Association("CreditCards").Append([]CreditCard{
		{Number: "0003"},
	})

}
