package main

import (
	"fmt"
)

// https://leetcode.cn/problems/reverse-string/

func reverseString(s []byte) {
	j := len(s)
	for i := 0; i < j/2; i++ {
		s[i], s[j-i-1] = s[j-i-1], s[i]
	}
	fmt.Println(string(s[:]))
}

func main() {
	str := "hello"
	s := []byte(str)
	reverseString(s)
}
