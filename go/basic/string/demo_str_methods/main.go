package main

import (
	"fmt"
	"strings"
)

func main() {
	// go 中的 str 类型其实是一个只读的数组，可以使用切片
	s := "hello world"
	s_zh := "你好世界"
	fmt.Printf("%v\n", s[:2])

	// 字符串的长度
	fmt.Println("######字符串的长度######")
	fmt.Println(len(s))
	fmt.Println(len(s_zh))

	// 字符串的拼接
	fmt.Println("######字符串的拼接######")
	s_new := s + s_zh
	fmt.Println(s_new)

	// 字符串的修改

	/*
		第一步:需要将字符串转成 rune 的 byte 数组 []rune(ss)
		第二步:切片获取原有数据 ss_rune[:2]
		第三步:string 方法将 byte 数组转换成字符串: string(ss_rune[:2])
	*/
	fmt.Println("######字符串的修改#####")
	ss := "你好啊"
	// ss_rune 不是一个字符串，已经变成了一个 rune 切片
	ss_rune := []rune(ss)
	fmt.Println(string(ss_rune))
	// string: 将其它数据类型转换成 string 类型
	fmt.Println(string(ss_rune[:2]))
	// 将 啊 修改成 吗
	s_world := string(ss_rune[:2]) + "吗"
	fmt.Println(s_world)

	// 将字符串转成数组: strings.Split
	fmt.Println("######将字符串转成数组#####")
	str := "11+12+13"
	arr1 := strings.Split(str, "+")
	fmt.Printf("%T,%v\n", arr1, arr1)

	// 将数组转换成字符串: strings.Join
	fmt.Println("######将数组转换成字符串#####")
	ss_new := strings.Join(arr1, "*")
	fmt.Printf("%T,%v\n", ss_new, ss_new)

	// 单引号只能表示一个字符
	fmt.Println("######单引号只能表示一个字符#####")
	s_a := 'a'
	// s_b := 'abc' // more than one character in rune literal
	fmt.Println(s_a)

}
