package main

import "fmt"

// 定义名为 Teacher 的结构体
type Teacher struct {
	Name   string
	Age    int
	School string
}

func main() {
	// 创建 Teacher 结构体的实例、对象、变量
	// 方式一:
	var t1 Teacher
	//输出未赋值的t1
	fmt.Println(t1)
	// 给t1赋值
	t1.Name = "苗总"
	t1.Age = 18
	t1.School = "新疆农大"
	// 输出赋值后的 t1
	fmt.Println(t1)
	// 输出赋值后的 t1.Age
	fmt.Println(t1.Age)
}
