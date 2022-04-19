package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//	第二步: 定义表结构(模型定义)
type User struct {
	// gorm:"primary_key" 主键索引,标记当前这个 Id 是自增的(1.2.3...)
	Id       int64 `json:"id" gorm:"primary_key"`
	Username string
	Password string
}

func main() {
	// 第一步: 连接数据库

	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:iPanel@202204121853@tcp(106.55.94.54:33333)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//fmt.Println(db)
	//fmt.Println(err)

	// 第三步: 创建表
	// 参考 https://gorm.io/zh_CN/docs/migration.html
	//db.AutoMigrate(User{})

	// 第四步： 增删改查
	// 单表
	// 4.1 增: 向 User 表添加一条数据
	//db.Create(&User{
	//	Username: "zhangsan",
	//	Password: "123456",
	//})

	//db.Create(&User{
	//	Username: "lisi",
	//	Password: "567890",
	//})

	// 4.2 改: 修改表的某一个字段
	// 将 Id 为1的记录 password修改为654321
	//db.Model(User{
	//	Id: 1,
	//}).Update("password", "654321")

	// 4.3 查询

	// 查询单条数据： First
	// SELECT * FROM `users` WHERE `users`.`id` = 1 ORDER BY `users`.`id` LIMIT 1
	u := User{Id: 1}
	db.First(&u)
	fmt.Printf("%#v \n", u)

	// 查询所有数据
	users := []User{} // 定义一个 User 结构体的切片来接收
	db.Find(&users)
	fmt.Printf("%v \n", users)

	// 4.4 删除

	// 根据主键删除
	// 删除 Id 为1的数据
	// DELETE FROM `users` WHERE `users`.`id` = 1
	db.Delete(&User{Id: 1})

	// 根据条件删除
	// DELETE FROM `users` WHERE username = 'lisi'
	db.Where("username = ?", "lisi").Delete(&User{})
}
