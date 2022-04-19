package main

import (
	"fmt"
	"time"
)

// defer + recover
// 问题处理
// 多个协程工作，其中一个协程出现 panic,导致程序崩溃
// 利用 defer+recover 捕获 panic 进行处理，即使协程出现问题，主线程仍然不受影响可以继续执行
// 输出数字:
func printNum() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

// 做除法操作
func devide() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("devide()出现错误：", err)
		}
	}()
	num1 := 10
	num2 := 0
	// 让 num1 / num2 报错
	result := num1 / num2
	fmt.Println(result)
}

func main() {
	// 启动两个协程:
	go printNum()
	go devide()
	time.Sleep(1 * time.Second)
}
