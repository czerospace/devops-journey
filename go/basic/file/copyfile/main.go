package main

import (
	"fmt"
	"io/ioutil"
)

// 复制文件 demo01.txt 中内容到 demo02.txt
func main() {
	// 定义源文件
	filepath01 := "demo01.txt"
	filepath02 := "demo02.txt"

	// 对文件进行读取
	context, err := ioutil.ReadFile(filepath01)
	if err != nil {
		fmt.Printf("读取文件出错: %v\n", err)
	}

	// 写入文件,文件不存在时默认创建
	err = ioutil.WriteFile(filepath02, context, 0666)
	if err != nil {
		fmt.Printf("写入文件出错: %v\n", err)
	}
}
