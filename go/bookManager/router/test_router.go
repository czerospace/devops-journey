package router

import "github.com/gin-gonic/gin"

func LoadTestRouter(r *gin.Engine) {
	r.GET("/test", func(c *gin.Context) {
		c.String(200, "for test")
	})
}
