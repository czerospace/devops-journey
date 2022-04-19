package main

import (
	"fmt"
	"strconv"
	"time"
)

// 在主线程中，开启一个 goroutine ，该 goroutine 每隔1s输出 "hello 苗总"
// 在主线程中，每隔1s输出 "hello 叶锅",输出10次后，退出程序
// 要求主线程和 goroutine 同时执行

func test() {
	for i := 1; i <= 10; i++ {
		fmt.Println("hello 苗总 + " + strconv.Itoa(i))
		// 阻塞一秒
		time.Sleep(time.Second)
	}
}

func main() {
	// 开启一个 goroutine
	go test()
	for i := 1; i <= 10; i++ {
		fmt.Println("hello 叶锅 + " + strconv.Itoa(i))
		// 阻塞一秒
		time.Sleep(time.Second)
	}
}
