package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 一对多关联查询 Association

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

	// 第三步: 查询关联表数据
	// 3.1 使用 Association 方法，需要把 User 查询好
	// 3.2 然后根据 User 定义指定的 AssociationForeignKey 去查找 CreditCard
	u := User{Username: "miaozong"}
	db.First(&u)
	fmt.Printf("%v \n", u)
	err := db.Model(&u).Association("CreditCards").Find(&u.CreditCards)
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println(u)
	// 结果转成 json 格式
	strUser, _ := json.Marshal(&u)
	fmt.Printf(string(strUser))
}
