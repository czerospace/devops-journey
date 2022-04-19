package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 中间件

// 第四步: 定义一个全局中间件（所有路由都会使用）
/*
MiddleWare 只是一个函数名字，可以随意定义
gin.HandlerFunc 中间件必须要返回的方法（如果不返回这个方法不是一个正确的中间件）
*/

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("我是一个全局中间件")
	}
}

func main() {
	// 第一步: 实例化引擎
	r := gin.Default()

	// 第二步: 路由
	r.GET("/hello", func(c *gin.Context) {
		fmt.Println("执行hello")
		c.JSON(200, gin.H{"msg": "success"})
	})
	// 第五步: 局部中间件调用
	r.GET("/part", MiddleWare(), func(c *gin.Context) {
		fmt.Println("执行part")
		c.JSON(200, gin.H{"msg": "success"})
	})

	// 第三步:启动
	fmt.Println("http://127.0.0.1:8000/hello")
	r.Run(":8000")
}
