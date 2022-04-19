package main

import "fmt"

func main() {
	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}
	sn3 := struct {
		name string
		age  int
	}{age: 11, name: "qq"}

	// mismatched types struct { age int; name string } and struct { name string; age int }
	if sn1 == sn3 {
		fmt.Println("sn1 == sn3")
	}

	sm1 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}
	sm2 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	// struct containing map[string]string cannot be compared)
	if sm1 == sm2 {
		fmt.Println("sm1 == sm2")
	}

}

// 1.结构体只能比较是否相等，但是不能比较大小
// 2.相同类型的结构体才能够进行比较，结构体是否相同不但与属性类型有关，还与属性顺序相关，sn3 与 sn1 就是不同的结构体
// 3.如果 struct 的所有成员都可以比较，则该 struct 就可以通过 == 或 != 进行比较是否相等，比较时逐个项进行比较，如果每一项都相等，则两个结构体才相等，否则不相等

// 总结：
// bool、数值型、字符、指针、数组等可以比较
// 切片、map、函数等是不能比较的。 具体可以参考 Go 说明文档。https://golang.org/ref/spec#Comparison_operators
