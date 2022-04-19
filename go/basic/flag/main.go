package main

import (
	"flag"
	"fmt"
)

func main() {
	var name string
	/*
			-name string
		        姓名 (default "zhangsan")
	*/
	flag.StringVar(&name, "name", "zhangsan", "姓名")
	/*
		go run main.go -name miaozong
	*/
	flag.Parse()
	// 获取指定参数
	fmt.Println(name)
	// 获取其他参数
	fmt.Println(flag.Args())
}
