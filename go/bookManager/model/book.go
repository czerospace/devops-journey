package model

/*
	定义 Book 表
*/

type Book struct {
	Id   int64  `json:"id" gorm:"primaryKey"`
	Name string `json:"name" binding:"required"`
	Desc string `json:"desc"`
	// 建立和 User表的多对多关系
	Users []User `gorm:"many2many:book_users"`
}
