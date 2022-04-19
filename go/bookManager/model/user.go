package model

/*
	定义 User 表
*/

type User struct {
	Id       int64  `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"not null" binding:"required"`
	Password string `json:"password" gorm:"not null" binding:"required"`
	Token    string `json:"token"`
}

// 表默认会添加 s ，自定义表的名字
func (User) tableName() string {
	return "user"
}
