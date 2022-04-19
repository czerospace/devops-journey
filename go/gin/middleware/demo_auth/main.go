package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 增加中间件
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 Header 中的 token
		token := c.Request.Header.Get("token")
		fmt.Println("get token is: ", token)
		// 如果 token 不是 twgdh 就返回 403
		if token != "twgdh" {
			c.String(http.StatusForbidden, "身份验证不通过")
			// 终止当前请求，不会将请求转发给路由，所以 处理函数不会执行
			c.Abort()
			return
		}
		// token 正确，向下执行 处理函数
		c.Next()
		fmt.Println("处理函数执行完成")
	}
}
func main() {
	r := gin.Default()

	// 首页，无需登陆直接访问
	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "index无需登陆直接访问"})
	})
	// home 页面，需要登陆
	r.GET("/home", Auth(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "home需要登陆验证token"})
	})

	r.Run(":8000")
}
