package main

import "fmt"

// new 和 make
// make 关键字的作用是创建于 slice、map 和 channel 等内置的数据结构
// new 的作用是为类型申请一片内存空间，并返回指向这片内存的指针,自定义类型使用 new 函数来分配空间
func main() {
	a := make([]int, 3, 10)
	a = append(a, 1)
	fmt.Printf("%v:%T \n", a, a)

	var b = new([]int)
	// b.append undefined (type *[]int has no field or method append
	// b = b.append(b, 2) // 返回的是内存指针，不能直接 append
	*b = append(*b, 3)
	fmt.Printf("%v:%T", b, b)
}
