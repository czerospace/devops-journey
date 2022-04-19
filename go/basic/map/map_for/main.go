package main

import "fmt"

// map 的遍历

func main() {
	userInfo := map[string]int{
		"zhangsan": 20,
		"lisi":     21,
		"wangwu":   23,
	}

	// 1.使用 for range 只遍历 key
	for k := range userInfo {
		if val, ok := userInfo[k]; ok {
			fmt.Println(k, val)
		}
	}

	// 2.遍历 key 和 value
	for key, value := range userInfo {
		fmt.Println(key, value)
	}
}
