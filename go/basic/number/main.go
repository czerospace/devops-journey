package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var a int8 = 4
	b := 4
	// var c int8 = 128 //constant 128 overflows int8
	fmt.Println(a, b)

	// 查看数据类型
	fmt.Println(reflect.TypeOf(a)) // int8
	fmt.Println(reflect.TypeOf(b)) // int
	fmt.Printf("b的类型是: %T\n", b)

	d := true
	fmt.Printf("d的类型是: %T\n", d)

	// 打印字节占用数 unsafe
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(b))
	fmt.Println(unsafe.Sizeof(d))
}
