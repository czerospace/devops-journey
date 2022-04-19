package main

import (
	"bookManager/dao/mysql"
	"bookManager/router"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化 mysql
	mysql.InitMysql()
	fmt.Println(mysql.DB, 111111111)
	// 1.将实例化 router 服务的方法拆分到 router 目录
	r := router.InitRouter()
	// 2.定义路由
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "ok")
	})
	// 3.启动服务

	r.Run(":8000")
}
