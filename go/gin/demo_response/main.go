package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// gin 中的 response

func main() {
	r := gin.Default()
	r.GET("/response", ResponseHandler)
	r.GET("/response/json", ResponseJsonHandler)
	r.GET("/response/json2", ResponseJson2Handler)
	r.GET("/response/redirect", ResponseRedirectHandler)
	r.Run(":8000")
}

// 1.响应一个普通的 String 字符串
// 一般测试时用一下
func ResponseHandler(c *gin.Context) {
	c.String(200, "响应一个string字符串")
}

// 2.返回一个 json 数据(最常用)
// 2.1通过结构体返回
func ResponseJsonHandler(c *gin.Context) {
	type Data struct {
		Msg  string `json:"msg"`
		Code int    `json:"code"`
	}

	// 假设从数据库查到的数据
	d := Data{
		Msg:  "Success",
		Code: 1001,
	}
	c.JSON(200, d)
}

// 2.2通过 gin.H 返回
func ResponseJson2Handler(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg":  "Success from gin.H",
		"code": 1001,
	})
}

// 3.路由重定向(基本不怎么用)
func ResponseRedirectHandler(c *gin.Context) {
	// 三秒后，重定向到百度
	time.Sleep(time.Second * 3)
	c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
}
