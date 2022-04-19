package router

import "github.com/gin-gonic/gin"

/*
	加载其他路由文件中的路由
*/

// InitRouter 这个方法作用: 初始化其他文件中的路由
func InitRouter() *gin.Engine {
	// 1.初始化 gin 服务
	r := gin.Default()
	// 加载  各个分层的路由
	LoadTestRouter(r)
	LoadAPIRouter(r)
	return r
}
