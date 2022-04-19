package main

import (
	"fmt"
)

// 接口的定义:定义规则、定义规范、定义某种能力
type SayHello interface {
	// 声明没有实现的方法:
	sayHello()
}

// 接口的实现:定义一个结构体:

// 中国人
type Chinese struct {
}

// 具体实现:
func (person Chinese) sayHello() {
	fmt.Println("你好")
}

// 美国人
type American struct {
}

// 具体实现:
func (person American) sayHello() {
	fmt.Println("hi")
}

// 定义一个函数: 专门用来各国人打招呼的函数，接收具备 SayHello 接口能力的变量
func greet(s SayHello) {
	s.sayHello()
}

func main() {
	// 创建一个中国人
	c := Chinese{}
	// 创建一个美国人
	a := American{}

	// 中国人打招呼
	greet(c)
	// 美国人打招呼
	greet(a)

	// 直接用接口创建实例,出错
	// var s SayHello
	// s.sayHello()

	// s 指向 c
	var s SayHello = c
	s.sayHello()
}
