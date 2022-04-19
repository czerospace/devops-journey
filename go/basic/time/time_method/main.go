package main

import (
	"fmt"
	"time"
)

// 时间的加减

func main() {
	now := time.Now()
	fmt.Println("当前时间:", now)

	// Add 方法
	// 获取10分钟前时间
	m, _ := time.ParseDuration("-1m")
	m1 := now.Add(10 * m)
	fmt.Println("十分钟前:", m1)
	// 获取8小时前时间
	h, _ := time.ParseDuration("-1h")
	h1 := now.Add(8 * h)
	fmt.Println("八小时前:", h1)
	// 获取1天前时间
	d, _ := time.ParseDuration("-24h")
	d1 := now.Add(d)
	fmt.Println("一天前:", d1)

	// Sub 方法
	// 求两个时间之间的差值
	subM := now.Sub(m1)
	fmt.Println(subM.Minutes(), "分钟")
	subH := now.Sub(h1)
	fmt.Println(subH.Hours(), "小时")
	subD := now.Sub(d1)
	fmt.Println(subD.Hours()/24, "天")

}
