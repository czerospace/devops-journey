package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1.创建路由 engine
	r := gin.Default()
	// 2.绑定路由
	/*
		"/"  是路由
		func(c *gin.Context) 是处理函数
	*/
	r.GET("/", func(c *gin.Context) {
		// 1.解析 get/post 请求的参数
		// 2.根据参数去操作数据库
		// 3.返回操作数据库的结果
		c.String(200, "hello world")
	})
	fmt.Println("server listen on http://127.0.0.1:8000")
	// 3.启动监听端口
	r.Run(":8000")
}
