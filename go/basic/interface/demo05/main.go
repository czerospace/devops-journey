package main

import "fmt"

// 定义两个接口
type Binterface interface {
	b()
}

type Cinterface interface {
	c()
}

// 定义A接口继承 B C接口
type Ainterface interface {
	a()
	Binterface
	Cinterface
}

// 定义一个结构体
type Stu struct {
}

// 实现 a() b() c()
func (s Stu) a() {
	fmt.Println("a")
}

func (s Stu) b() {
	fmt.Println("b")
}

func (s Stu) c() {
	fmt.Println("c")
}

// 定义一个空接口
type E interface {
}

func main() {
	var s Stu
	var e E = s
	fmt.Println(e)
	var e2 interface{} = s
	fmt.Println(e2)
	var num float64 = 9.3
	var e3 interface{} = num
	fmt.Println(e3)
}
