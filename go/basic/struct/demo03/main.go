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
	// 方式三:
	var t3 *Teacher = new(Teacher)
	// t3 是指针，t3 其实指向的是一个地址，应该给这个地址 指向的对象的字段赋值
	(*t3).Name = "苗总"
	(*t3).Age = 18
	// 为了符合程序员的编程习惯，go 提供了简化的赋值方法
	t3.School = "新疆农大" // go 编译器底层对 t3.School 转化为 (*t3).School="新疆农大"
	fmt.Println(*t3)
}
