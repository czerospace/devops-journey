package main

import (
	"fmt"
	"sync"
)

// 定义一个 wg 无需赋值
var wg sync.WaitGroup

func main() {
	// 启动五个协程
	for i := 1; i <= 5; i++ {
		// 协程开始的时候加1操作
		wg.Add(1)
		go func(n int) {
			// 协程执行完成减1，防止忘记减1操作，结合 defer 使用
			defer wg.Done()
			fmt.Println(n)
		}(i)

	}

	// 主线程一直在阻塞，什么时候 wg 减为0了，就停止
	wg.Wait()
}
