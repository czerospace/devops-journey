package main

import (
	"fmt"
)

// select

func main() {
	// 定义多个管道
	intChan := make(chan int, 1)
	stringChan := make(chan string, 1)

	go func() {
		// time.Sleep(1 * time.Second)
		intChan <- 10
	}()

	go func() {
		// time.Sleep(2 * time.Second)
		stringChan <- "mznb"
	}()

	select {
	case v := <-intChan:
		fmt.Println("intChan: ", v)
	case v := <-stringChan:
		fmt.Println("stringChan: ", v)
	default:
		fmt.Println("防止 select 被阻塞")
	}
}
