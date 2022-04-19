package main

import (
	"fmt"
	"os"
)

func main() {
	// 打开文件
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Printf("文件打开出错，对应的错误为: %v\n", err)
	}
	fmt.Printf("文件=%v\n", file)

	// 关闭文件
	err = file.Close()
	if err != nil {
		fmt.Printf("文件关闭失败: %v\n", err)
	}
}
