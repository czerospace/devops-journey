package controller

import (
	"bookManager/dao/mysql"
	"bookManager/model"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

// RegisterHandler 注册功能
func RegisterHandler(c *gin.Context) {
	p := new(model.User)
	// 参数校验，绑定
	if err := c.ShouldBind(p); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}
	// 入库: 最简单的创建数据
	mysql.DB.Create(p)
	c.JSON(200, gin.H{"msg": "success"})
}

// LoginHandler 登陆功能
func LoginHandler(c *gin.Context) {
	p := new(model.User)
	// 参数校验
	if err := c.ShouldBind(p); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
	}
	// 判断用户的用户名和密码是否正确
	u := model.User{Username: p.Username, Password: p.Password}
	// rows 判断是否有数据 nil 证明没有这个用户
	if rows := mysql.DB.Where(&u).First(&u).Row(); rows == nil {
		c.JSON(403, gin.H{"msg": "密码错误"})
		return
	}
	// 随机生成一个字符串作为 token
	token := uuid.New().String()
	// 将 token 写入数据库
	mysql.DB.Model(u).Update("token", token)
	c.JSON(200, gin.H{"msg": "success", "token": token})
}
