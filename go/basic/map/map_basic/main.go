package main

import "fmt"

// map 基础
// map 是一种  无序 的 key-value 数据结构

func main() {
	// 1.定义 map
	userInfo := map[string]int{
		"zhangsan": 24,
		"lisi":     28,
	}
	fmt.Println(userInfo)

	// 2.判断一个 key 是否在 map 中
	val, ok := userInfo["zhangsan"]
	fmt.Println(val, ok) // 24 true
	val, ok = userInfo["tony"]
	fmt.Println(val, ok) // 0 false

	// 3.和 if 一起常用写法
	if _, ok := userInfo["zhangsan"]; ok {
		fmt.Println(userInfo["zhangsan"])
	} else {
		fmt.Println("map中不存在这个key")
	}

	// 4.删除 map 中的一个数据
	fmt.Println(userInfo)
	delete(userInfo, "lisi")

	if _, ok := userInfo["lisi"]; ok {
		fmt.Println(userInfo["lisi"])
	} else {
		fmt.Println("map中不存在这个key")
	}
}
