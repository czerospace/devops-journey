package main

import "fmt"

// 结构体之间的转换二

type Student struct {
	Age int
}

type Stu Student

func main() {
	// 结构体进行 type 重新定义(相当于取别名)，Golang 认为是新的数据类型，但是相互间可以强转
	var s1 Student = Student{18}
	var s2 Stu = Stu{19}
	s1 = Student(s2)
	fmt.Println(s2)
	fmt.Println(s1)
}
