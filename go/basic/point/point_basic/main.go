package main

import "fmt"

// 指针

// 指针地址(&a)
// 指针取值(*&a)
// 指针类型(&a)

func main() {
	var a = 10
	fmt.Printf("指针地址: %d \n", &a)
	fmt.Printf("指针取值: %d \n", *&a)
	fmt.Printf("指针类型: %T \n", &a)
}
