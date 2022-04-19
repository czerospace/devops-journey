package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

/*
	使用 net/http 模块作为 web 的客户端发送 get 请求
	应用场景:
		1) 爬虫，获取页面数据
		2) 调用其他服务中的接口
*/

func main() {
	// 1.直接通过 Url 拼接出 Url 字符串
	apiUrl := "http://127.0.0.1:8080/req/get?name=miaozong"
	// apiUrl := "https://www.baidu.com"

	resp, err := http.Get(apiUrl)
	if err != nil {
		fmt.Println(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	// 2.通过 net/url 进行解析
	apiUrl2 := "http://127.0.0.1:8080/req/get"
	data := url.Values{}
	data.Set("name", "miaozong")
	u, _ := url.ParseRequestURI(apiUrl2)
	// Encode 一下数据
	u.RawQuery = data.Encode()
	fmt.Println(u.String()) // http://127.0.0.1:8080/req/get?name=miaozong

	resp2, err2 := http.Get(u.String())
	if err2 != nil {
		fmt.Println(err2)
	}
	body2, _ := ioutil.ReadAll(resp2.Body)
	fmt.Println(string(body2))
}
