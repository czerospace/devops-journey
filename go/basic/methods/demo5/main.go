package main

import "fmt"

// 方法的注意事项

// 4.方法的访问范围的控制规则和函数一样。方法名首字母小写，只能在本包访问，方法名首字母大写，可以在本包和其他包访问。

// 5.如果一个类型实现了 String ，那么 fmt.Println 默认会调用这个变量的 String() 进行输出

type Student struct {
	Name string
	Age  int
}

// 定义结构体的时候，常定义 String() 作为输出结构体信息的方法，当fmt.Println会自动调用

func (s *Student) String() string {
	str := fmt.Sprintf("Name = %v , Age = %v", s.Name, s.Age)
	return str
}

func main() {
	stu := Student{
		Name: "苗总",
		Age:  18,
	}
	// 传入地址，如果绑定了 String 方法就会自动调用
	fmt.Println(&stu)
}
