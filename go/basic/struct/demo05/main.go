package main

import "fmt"

// 结构体之间的转换一

type Student struct {
	Age int
}

type Person struct {
	Age int
}

func main() {
	// 结构体是用户单独定义的类型，和其它类型进行转换时需要有完全相同的字段(名字、个数和类型)
	var s Student = Student{10}
	var p Person = Person{9}
	s = Student(p)
	fmt.Println(s) //9
	fmt.Println(p) //9
}
