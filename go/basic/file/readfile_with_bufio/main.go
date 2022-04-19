package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 使用带缓冲的方式按行读取文件
// 适用于大文件
func main() {
	// 打开文件
	file, err := os.Open("bufio.txt")
	if err != nil {
		fmt.Printf("文件打开失败: %v\n", err)
	}
	// 当函数退出时，关闭 file，防止内存泄露
	defer file.Close()

	// 创建一个流
	reader := bufio.NewReader(file)
	// 读取流
	for {
		// 读取到一个换行就结束
		str, err := reader.ReadString('\n')
		// 如果没有读取到文件结尾的话，就正常输出文件内容即可
		fmt.Println(str)
		// io.EOF 表示已经读取到文件的结尾
		if err == io.EOF {
			break
		}
	}

	// 结束
	fmt.Println("文件读取结束")
}
