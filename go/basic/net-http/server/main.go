package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
使用 net/http 模块，开发一个简单的 web-server
	1.提供 get 请求
	2.提供 post 请求
http: 协议(和 web 服务器进行交互的规范，规则)
	Get: 从数据库读取数据，比如 查询订单
		- get 请求参数是从 url 上直接读取(一般数据都比较少)
	Post: 创建新的数据，比如 12306买票(在数据库添加一条数据，购票记录)
		- post 是从 http 的 body 中获取数据
	Put: 修改数据，比如 更新支付宝上的用户信息
	Delete: 删除数据库中的数据
*/

// 开发一个 web 服务器主要的步骤
/*
	第一步: 路由
	第二步: 处理数据
	- 解析请求的数据(比如获取某一个商品，需要把商品 Id 信息携带给后端)
		- 根据请求数据参数查询数据库
	- 响应数据(把从数据库读取的数据，返回给浏览器或者请求方)
*/

func main() {
	// 第一步: 路由

	/*
		"/req/get": 路由 指 URL中除去域名的那一块(http://www.example.com/req/get)
		dealGetHandler: 处理函数(处理服务请求)
	*/
	// http://127.0.0.1:8080/req/get?name=miaozong
	http.HandleFunc("/req/get", dealGetHandler)

	http.HandleFunc("/req/post", dealPostHandler)

	fmt.Println("server starting")

	// 第三步: 启动服务
	/*
		addr:当前 server 监听的ip和端口
		handler: 处理函数
	*/
	http.ListenAndServe(":8080", nil)

}

// 处理 get 请求
/*
	1).解析请求的数据(比如获取某一个商品，需要把商品 Id 信息携带给后端)
	http.Request： 解析 url 中的数据或者 post 请求中 body 的数据
	2).响应数据(把从数据库读取的数据，返回给浏览器或者请求方)
	http.ResponseWriter: 本质是一个 interface 接口,定义了三个方法，返回数据
*/

// 处理函数的名字用驼峰命名: xxxHandler 函数名
func dealGetHandler(w http.ResponseWriter, r *http.Request) {
	// 1).解析请求的数据
	query := r.URL.Query() // 返回 map[string][]string
	// 1.1 通过字典下标取路由参数
	if len(query["name"]) > 0 {
		names := query["name"][0]
		fmt.Println("字典下标取值", names)
	}

	// 1.2 通过 Get 方法取值
	name2 := query.Get("name")
	fmt.Println("通过get方法取值", name2)
	fmt.Println(query)

	// 2).响应数据
	// 2.1 返回一个简单字符串
	// w.Write([]byte("hello world!"))

	// 2.2 返回一个 json 数据
	// 假设我们拿到了 name=miaozong 我们到数据库取出了 miaozong 用户的信息
	type Info struct {
		Name     string
		Password string
		Age      int
	}
	// 假设这是从数据库中取出的
	u := Info{
		Name:     name2,
		Password: "123456",
		Age:      24,
	}
	json.NewEncoder(w).Encode(u)
}

// Info 结构体，包含用户名和用户密码
type Info struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// 和 get 请求是一模一样的写法，这是一个 post 请求
func dealPostHandler(w http.ResponseWriter, r *http.Request) {
	// 1).解析请求的数据
	// r.URL.Query() 从 url 取请求参数
	// post 请求从 http 的 body 中获取数据
	bodyContent, _ := ioutil.ReadAll(r.Body) // 返回的是一个 byte
	//fmt.Printf("%T %v", bodyContent, bodyContent)

	// 获取 string
	// strData := string(bodyContent)
	// 如何才能解析这个 string 字符串(string 转结构体)
	// 先定义一个字段一模一样的结构体

	var d Info
	// json.Unmarshal([]byte(strData), &d)
	json.Unmarshal(bodyContent, &d) // 使用 Unmarshal 将 bodyContent 解析为结构体
	fmt.Println("获取的数据name:", d.Name)
	fmt.Println("获取的数据password", d.Password)
	// 2).响应数据
	// 2.1 返回一个简单字符串
	w.Write([]byte("hello miaozong nb"))
}
