package main

import "fmt"

// 管道的关闭

func main() {
	var intChan chan int
	intChan = make(chan int, 3)

	// 在管道中存放数据:
	intChan <- 10
	intChan <- 20

	// 关闭管道
	close(intChan)

	// 再次写入数据：panic: send on closed channel
	// 管道关闭后不能写入数据
	// intChan <- 30

	// 管道关闭后可以读取数据
	num := <-intChan
	fmt.Println(num)
}
