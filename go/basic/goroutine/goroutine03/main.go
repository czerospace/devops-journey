package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	// 启动一个 goroutine
	// 使用匿名函数，直接调用匿名函数
	go func() {
		fmt.Println("匿名函数1")
	}()

	for i := 1; i <= 5; i++ {
		// 将外部变量 i 传参给匿名函数
		go func(n int) {
			fmt.Println("匿名函数2 + " + strconv.Itoa(n))
		}(i)
	}
	time.Sleep(2 * time.Second)
}
