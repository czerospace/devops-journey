package main

import "fmt"

// 创建 Teacher 结构体的实例、对象、变量
type Teacher struct {
	Name   string
	Age    int
	School string
}

func main() {
	// 创建 Teacher 结构体的实例、对象、变量
	// 方式二:
	var t2 Teacher = Teacher{"苗总", 18, "新疆农大"}
	fmt.Println(t2)
}
