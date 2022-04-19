package main

import (
	"fmt"
	"time"
)

func main() {
	// 1.时间对象
	now := time.Now()
	fmt.Printf("%T %v \n", now, now)

	// 2.格式化时间
	strTime := now.Format("2006-01-02 15:04:05")
	fmt.Printf("%T %v \n", strTime, strTime)

	// 3.时间戳
	ts := now.Unix()
	fmt.Printf("%T %v \n", ts, ts)

	// 4.格式化时间转成时间对象
	// 设置时区
	loc, _ := time.LoadLocation("Asia/Shanghai")
	// 转换成对应时区的时间对象
	timeObj, _ := time.ParseInLocation("2006-01-02 15:04:05", strTime, loc)
	fmt.Printf("%T %v \n", timeObj, timeObj)
	fmt.Println(timeObj.Unix()) // 时间戳
}
