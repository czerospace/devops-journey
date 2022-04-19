package main

import "fmt"

// 方法的引入

// 1.方法作用在指定的数据类型上、和指定的数据类型绑定，因此自定义类型都可以有方法，而不仅仅是struct

// 2.方法的声明和调用

// 定义 Person 结构体
type Person struct {
	Name string
}

// 给 Person 结构体绑定方法 test , 输出 Name
// 参数名字随便起，一般是结构体名称的首字母
func (p Person) test() {
	fmt.Println(p.Name)
}

func main() {
	// 创建结构体对象：
	var someone Person
	someone.Name = "苗总"
	// 调用 test 方法输出 Name 值
	someone.test()
}
