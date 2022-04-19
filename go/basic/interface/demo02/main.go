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

// 自定义数据类型
type integer int

func (i integer) sayHello() {
	fmt.Println("say hi + ", i)
}

func main() {
	var i integer = 10
	var s SayHello = i
	s.sayHello() // say hi +  10
}
