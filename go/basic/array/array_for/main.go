package main

import "fmt"

// 遍历数组

func main() {
	a := [...]string{"bj", "sh", "gz"}

	// 1.for 遍历
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	// 2.for range 遍历
	for index, value := range a {
		fmt.Println(index, value)
	}
}
