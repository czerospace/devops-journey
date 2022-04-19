package main

import (
	"fmt"
	"sync"
)

// 互斥锁

var totalNum int
var wg sync.WaitGroup

// 加入互斥锁
var lock sync.Mutex

func add() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		lock.Lock()
		totalNum = totalNum + 1
		lock.Unlock()
	}
}

func sub() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		lock.Lock()
		totalNum = totalNum - 1
		lock.Unlock()
	}
}

func main() {
	wg.Add(2)
	// 启动协程
	go add()
	go sub()
	wg.Wait()
	// 最终结果一定为0
	fmt.Println(totalNum)
}
