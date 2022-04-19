package routers

import "github.com/gin-gonic/gin"

func LoadBooks(e *gin.Engine) {
	e.GET("/book", BookHandler)
}

func BookHandler(c *gin.Context) {
	c.String(200, "书籍模块")
}
