package main

import (
	"fmt"
	"os"
)

func main() {
	// 1.获取目录
	fmt.Println(os.Getwd())

	// 2.切换路径
	// os.Chdir("/etc")
	// fmt.Println(os.Getwd())

	// 3.创建文件夹
	os.Mkdir("test_dir", 0777)

	// 4.删除文件夹
	os.Remove("test_dir")

}
