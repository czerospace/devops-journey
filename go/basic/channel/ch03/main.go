package main

import "fmt"

// 管道的遍历

func main() {
	var intChan chan int
	intChan = make(chan int, 100)

	// 给管道内写入100个数据
	for i := 0; i < 100; i++ {
		intChan <- i
	}

	// 遍历:在遍历前一定要关闭管道，不然会 deadlock
	close(intChan)
	for v := range intChan {
		fmt.Println("value = ", v)
	}
}
