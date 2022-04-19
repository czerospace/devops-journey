package main

import "fmt"

// 定义两个接口
type Ainterface interface {
	a()
}

type Binterface interface {
	b()
}

// 定义一个结构体
type Stu struct {
}

// 用 Stu 分别实现 A B 接口
func (s Stu) a() {
	fmt.Println("实现了A接口")
}

func (s Stu) b() {
	fmt.Println("实现了B接口")
}

func main() {
	var s Stu
	var a Ainterface = s
	var b Binterface = s
	a.a()
	b.b()
}
