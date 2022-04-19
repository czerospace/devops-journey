package main

import "fmt"

// 数组特点
//	1.长度固定(len)
//	2.连续内存空间
//	3.同一类型集合

// 数组只有长度，没有容量

func main() {
	// 1.定义一个数组
	var a [5]int
	fmt.Println(a) // [0 0 0 0 0] 不是 []

	// 2.向数组中指定索引添加一个元素
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	// 3 ... 让编译器自动识别数组的长度
	var b = [...]string{"bj", "sh", "gz"}
	fmt.Println(b)
}
