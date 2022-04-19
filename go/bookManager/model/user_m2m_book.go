package model

/*
	作用:用户与书籍管理的关系表
*/

type BookUser struct {
	UserId int64 `gorm:"primaryKey"`
	BookID int64 `gorm:"primaryKey"`
}
