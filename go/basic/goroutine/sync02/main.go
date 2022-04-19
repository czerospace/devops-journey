package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// 已知协程次数的时候可以先 wg.Add()
	// wg.Add()中加入的数字和协程的次数一定要保持一致
	wg.Add(5)
	for i := 1; i <= 5; i++ {
		go func(n int) {
			defer wg.Done()
			fmt.Println(n)
		}(i)
	}

	wg.Wait()
}
