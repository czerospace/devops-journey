package main

import (
	"fmt"
	"sync"
	"time"
)

// 协程和管道协同工作
// 1. 开启一个 writeData 协程，向管道中写入50个整数
// 2. 开启一个 readData 协程，从管道中读取 writeDate 写入的数据
// 3. 注意: wirteData 和 readData 操作的是同一个管道
// 4. 主线程需要等待 writeData 和 readData 协程都完成工作才能退出

var wg sync.WaitGroup

// 写
func writeData(intChan chan int) {
	defer wg.Done()
	for i := 1; i <= 50; i++ {
		intChan <- i
		fmt.Println("写入的数据为: ", i)
		time.Sleep(time.Second)
	}
	// 管道关闭
	close(intChan)
}

// 读
func readData(intChan chan int) {
	defer wg.Done()
	// 遍历:
	for v := range intChan {
		fmt.Println("读取的数据为:", v)
		time.Sleep(time.Second)
	}
}

func main() {
	intChan := make(chan int, 50)

	wg.Add(2)
	go writeData(intChan)
	go readData(intChan)
	wg.Wait()
}
