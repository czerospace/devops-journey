package main

import "fmt"

// 方法的注意事项

// 3.Golang 中的方法作用在指定的数据类型上，和指定的数据类型绑定，因此自定义类型都可以有方法，而不仅仅是 struct，比如 int，float32等都可以有方法

// 自定义一个 integer 类型，相当于 int 的别名
type integer int

// 创建一个 print 方法
func (i integer) print() {
	i = 30
	fmt.Println("i = ", i)
}

// 创建一个 change 方法
func (i *integer) change() {
	*i = 40
	fmt.Println("i = ", *i)
}

func main() {
	var num integer
	num = 20
	fmt.Println("num is ", num) // num is  20
	num.print()                 // i =  30
	num.change()                // i =  40
}
