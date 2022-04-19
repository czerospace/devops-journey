package main

import "fmt"

// 方法的注意事项

// 1.结构体类型是值类型，在方法调用中，遵守值类型的传递机制，是值拷贝传递方式

// 2.如果程序员希望在方法中，更改结构体变量的值，可以通过结构体指针的方式来处理

type Person struct {
	Name string
}

func (p *Person) test() {
	p.Name = "叶锅"       // 原始写法 (*p).name = "叶锅"
	fmt.Println(p.Name) // 原始写法 fmt.Println((*p).Name)
}

// 底层编译器做了优化，底层会自动帮我们加上 & *

func main() {
	var someone Person
	someone.Name = "苗总"
	fmt.Println(someone.Name) // 苗总
	someone.test()            // 叶锅 原始写法 (&someone).test()
	// 注意 test 方法中传的是指针，改变了内存地址里的值
	fmt.Println(someone.Name) // 叶锅
}
