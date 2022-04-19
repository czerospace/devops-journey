package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 四种注册路由方式

func main() {
	// 1.生成 engine
	r := gin.Default()

	// 2.注册路由

	// 2.1 无参数路由
	r.GET("/hello", HelloHandler)
	// 2.2 API 路由: http://127.0.0.1:8000/book/9527
	// :id 使用 id 来获取 9527 的值
	r.GET("/book/:id", GetBookDetailHandler)
	// 2.3 url 传参: http:/127.0.0.1:8000/user?id=20&name=zhangsan
	r.GET("/user", GetUserDetail)
	// 2.4 shouldBind 绑定(解析 post 请求中复杂的 json 数据)
	r.POST("/login/", LoginHandler)

	// 3.启动服务
	r.Run(":8000")
}

// 2.1 无参数路由
func HelloHandler(c *gin.Context) {
	c.String(200, "hello")
}

// 2.2 API 路由
func GetBookDetailHandler(c *gin.Context) {
	bookId := c.Param("id")
	fmt.Println("获取的bookID---->", bookId)
	c.String(200, "API params")
}

// 2.3 url 传参
func GetUserDetail(c *gin.Context) {
	// 1.获取值，如果没有为 nil
	name := c.Query("name")
	// 2.获取值，如果没有使用默认值
	name2 := c.DefaultQuery("name", "default value")
	fmt.Println("获取的用户名---->", name, name2)
	c.String(200, "URL params")
}

// 2.4 shouldBind 绑定

// shouldBind 方法获取 json 中复杂数据
type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func LoginHandler(c *gin.Context) {
	var login Login

	// c.ShouldBind(&login) 方法必须要传入一个结构体对象
	// 将 net/http 中的 r.Body 数据解析到 Login 结构体中
	if err := c.ShouldBind(&login); err != nil {
		c.String(200, "参数校验错误")
	}
	fmt.Println(login.Username, login.Password)
	c.String(200, "success")
}
