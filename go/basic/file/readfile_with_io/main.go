package main

import (
	"fmt"
	"io/ioutil"
)

// 读取文件,适用于一次性读取小文件

func main() {
	// ioutil.ReadFile 已封装 open close
	content, err := ioutil.ReadFile("io.txt") // 返回内容为: []byte, error

	if err != nil {
		fmt.Printf("读取文件出错: %v", err)
	}

	fmt.Printf("%v", string(content))
}
