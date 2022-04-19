package main

import "fmt"

// 只写 只读 管道

func main() {
	// 默认情况下，管道是双向的: 可读可写
	// var intChan1 chan int
	// 声明为只写:
	// 管道具备 <- 只写性质
	var intChan2 chan<- int
	intChan2 = make(chan int, 3)
	intChan2 <- 10
	fmt.Println("intChan2: ", intChan2)

	// 如果读管道 intChan2 会报错
	// num := <-intChan2 // invalid operation: <-intChan2 (receive from send-only type chan<- int)

	// 声明为只读:
	var intChan3 <-chan int
	if intChan3 != nil {
		num1 := <-intChan3
		fmt.Println("num1: ", num1)
	}

	// 往 intChan3 里面写数据
	// intChan3 <- 20 // invalid operation: intChan3 <- 20 (send to receive-only type <-chan int)
}
