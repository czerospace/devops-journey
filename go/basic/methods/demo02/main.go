package main

import "fmt"

// 方法的引入

// 3.注意
// test 方法种参数的名字随意起
// 结构体 Person 和 test 方法绑定，调用 test 方法必须靠指定的类型: Person
// 如果其它类型变量调用 test 方法一定会报错
// 结构体对象传入 test 方法中，值传递，和函数参数传递一致。

// 定义 Person 结构体
type Person struct {
	Name string
}

func (p Person) test() {
	p.Name = "叶锅"
	fmt.Println(p.Name)
}

func main() {

	var someone Person
	someone.Name = "苗总"

	someone.test()            // 输出叶锅
	fmt.Println(someone.Name) // 值传递，someone.Name 输出苗总
}
