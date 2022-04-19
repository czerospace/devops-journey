package main

import "fmt"

// 切片特点
//	1.长度不固定，可变(len)
//	2.连续内存空间
//	3.同一类型集合

// 切片有长度，有容量

func main() {
	// 1.定义切片
	num := []int{1, 2, 3, 4, 5}
	strs := []string{"zhangsan", "lisi"}
	fmt.Println(num, strs)
	fmt.Printf("长度: %v,容量: %v\n", len(num), cap(num))
}
