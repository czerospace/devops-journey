package main

import "fmt"

// 字符串的遍历
func main() {
	s := "hello 张三"

	// 1.for 遍历字符串
	// 不会按字符打印出
	for i := 0; i < len(s); i++ {
		fmt.Printf("%T:%v\n", s[i], s[i])
	}

	// 2.使用 range 遍历字符串
	for index, val := range s {
		fmt.Println(index, val)
		fmt.Printf("%T:%v:%c\n", val, val, val)
	}

}
