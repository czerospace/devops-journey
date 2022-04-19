package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	url := "http://127.0.0.1:8080/req/post"
	// 模拟 form 表单提交数据 contentType := "application/x-www-form-urlencoded"
	// 传json数据 json contentType := "application/json"
	contentType := "application/json"
	data := `{
		"name":"miaozong",
		"password":"123456"
	}`

	resp, _ := http.Post(url, contentType, strings.NewReader(data))
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))
}
