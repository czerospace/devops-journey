package main

import "fmt"

// 方法和函数的区别

// 1.绑定指定类型：
// 方法：需要绑定指定数据类型
// 函数：不需要绑定数据类型

type Student struct {
	Name string
}

// 定义方法
func (s Student) method01() {
	fmt.Println(s.Name)
}

// 定义函数
func function01(s Student) {
	fmt.Println(s.Name)
}

// 2.调用方式不一样：

// 函数的调用方式：
// 函数名(实参列表)

// 方法的调用方式：
// 变量 方法名(实参列表)

// 3.对于函数来说，参数类型对应的是什么就要传入什么

func function02(s *Student) {
	fmt.Println(s.Name)
}

// 4.对于方法来说，接收者为值类型，可以传入指针类型，接收者为指针类型，可以传入值类型

func (s *Student) method02() {
	fmt.Println((*s).Name)
}

func main() {
	var stu Student = Student{"zhangsan"}
	// 调用函数
	function01(stu)
	// 调用方法
	stu.method01()

	fmt.Println("--------分割线---------")

	// 函数传入值和指针
	// function02(stu) cannot use stu (type Student) as type *Student in argument to function02
	function02(&stu) // 只能传入 &stu
	// 方法传入值和指针
	stu.method01()
	(&stu).method01() // 虽然用指针类型调用，但是传递还是按照值传递的形式

	fmt.Println("--------分割线---------")
	// &stu 等价于 stu
	(&stu).method02()
	stu.method02()
}
