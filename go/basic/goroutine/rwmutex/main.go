package main

import (
	"fmt"
	"sync"
	"time"
)

// 读写锁
// 加入读写锁，默认是写锁
var lock sync.RWMutex

var wg sync.WaitGroup

func read() {
	defer wg.Done()
	lock.RLock()
	fmt.Println("开始读取数据")
	time.Sleep(time.Second)
	fmt.Println("读取数据成功")
	lock.RUnlock()
}

func write() {
	defer wg.Done()
	lock.Lock()
	fmt.Println("开始写入数据")
	time.Sleep(5 * time.Second)
	fmt.Println("写入数据成功")
	lock.Unlock()
}

func main() {
	wg.Add(6)
	// 启动协程 ---> 读多写少，读五次，写一次
	for i := 0; i < 5; i++ {
		go read()
	}
	go write()
	wg.Wait()
}
