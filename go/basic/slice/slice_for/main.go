package main

import "fmt"

// slice 的遍历

func main() {
	a := []string{"bj", "sh", "gz", "sz"}

	// 1. for 遍历
	for i := 0; i < len(a); i++ {
		fmt.Println(i, a[i])
	}

	// 2. for range 遍历
	for index, value := range a {
		fmt.Println(index, value)
	}
}
