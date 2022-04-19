package main

import (
	"demo_router_layer/routers"
	"github.com/gin-gonic/gin"
)

// 路由分层

func main() {
	// 1.创建路由engine
	r := gin.Default()

	// 2.注册路由
	routers.LoadUsers(r)
	routers.LoadBooks(r)
	// 3.启动服务
	r.Run(":8000")
}
