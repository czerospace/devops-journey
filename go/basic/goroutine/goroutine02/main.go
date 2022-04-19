package main

import (
	"fmt"
	"strconv"
	"time"
)

// 主死从随
// 主线程退出，goroutine 也会退出
func test() {
	for i := 1; i <= 1000; i++ {
		fmt.Println("我是 协程 + " + strconv.Itoa(i))
		// 阻塞一秒
		time.Sleep(time.Second)
	}
}

func main() {
	// 开启一个 goroutine
	go test()
	for i := 1; i <= 10; i++ {
		fmt.Println("我是主线程 + " + strconv.Itoa(i))
		// 阻塞一秒
		time.Sleep(time.Second)
	}
}
