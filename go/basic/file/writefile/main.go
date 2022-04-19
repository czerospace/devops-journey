package main

import (
	"bufio"
	"fmt"
	"os"
)

// 文件的写入
func main() {
	// 打开文件
	file, err := os.OpenFile("demo.txt", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("打开文件失败: %v", err)
	}
	// 将文件关闭
	defer file.Close()

	// 写入文件操作: --->> IO 流 --->> 缓冲输出流(带缓冲区)
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("苗总nb\n")
	}

	// 刷新缓冲区中的数据--->真正写入文件中
	writer.Flush()
}
