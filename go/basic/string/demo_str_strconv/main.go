package main

import (
	"fmt"
	"strconv"
)

// 字符串 转换

func main() {
	// 1.将 int 转换成字符串
	num := 100
	fmt.Printf("%T %d \n", num, num)
	strNum := strconv.Itoa(num)
	fmt.Printf("%T %v \n", strNum, strNum)

	// 2.将字符串转 int
	intNum, err := strconv.Atoi(strNum)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T:%v\n", intNum, intNum)
}
