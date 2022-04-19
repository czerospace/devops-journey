package main

import "fmt"

// 管道的入门

func main() {
	// 定义管道、声明管道
	// 定义一个 int 类型的管道
	var intChan chan int
	// 通过 make 初始化:管道存放3个 int 类型的数据
	intChan = make(chan int, 3)

	// 证明管道是引用类型:
	fmt.Printf("intChan 的值: %v", intChan) // 0xc000110080

	// 向管道存放数据:
	intChan <- 10
	num := 20
	intChan <- num
	intChan <- 30
	// 注意: 不能存放大于容量的数据
	// intChan <- 40 //error: all goroutines are asleep - deadlock!

	// 从管道中读取数据:
	num1 := <-intChan
	num2 := <-intChan
	num3 := <-intChan
	fmt.Println(intChan)
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)

	// 注意:在没有使用协程的情况下，如果管道的数据已经全部取出，那么再取就会报错.
	// num4 := <-intChan
	// fmt.Println(num4) // fatal error: all goroutines are asleep - deadlock!

	// 输出管道的长度
	fmt.Printf("管道的实际长度是: %v,管道的容量是: %v", len(intChan), cap(intChan))
}
