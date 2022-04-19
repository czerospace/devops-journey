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
	name string
}

// 具体实现:
func (person Chinese) sayHello() {
	fmt.Println("你好")
}

// 中国人特有的方法
func (person Chinese) niuYangGe() {
	fmt.Println("东北文化-扭秧歌")
}

// 美国人
type American struct {
	name string
}

// 具体实现:
func (person American) sayHello() {
	fmt.Println("hi")
}

func (person American) disco() {
	fmt.Println("disco")
}

// 定义一个函数: 专门用来各国人打招呼的函数，接收具备 SayHello 接口能力的变量
func greet(s SayHello) {
	s.sayHello()
	// 断言：看s是否能转成 Chinese 类型并且赋给 ch 值,flag是判断是否成功
	// ch, flag := s.(Chinese)
	// if flag == true {
	// 	// 只有中国人有 niuYangGe()
	// 	ch.niuYangGe()
	// } else {
	// 	fmt.Println("美国人不会扭秧歌")
	// }

	// if ch, ok := s.(Chinese); ok {
	// 	ch.niuYangGe()
	// } else {
	// 	fmt.Println("外国人不会扭秧歌")
	// }

	// type 属于 go 中的一个关键字，固定写法
	switch s.(type) {
	case Chinese:
		ch := s.(Chinese)
		ch.niuYangGe()
	case American:
		am := s.(American)
		am.disco()
	}

	fmt.Println("打招呼")
}

func main() {
	// 创建一个中国人
	c := Chinese{}

	// 中国人打招呼
	greet(c)

	// 创建一个美国人
	a := American{}
	greet(a)
}
