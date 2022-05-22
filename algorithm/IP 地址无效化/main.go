package main

import (
	"fmt"
	"strings"
)

// https://leetcode.cn/problems/defanging-an-ip-address/
// 思路就是遍历字符串，然后把 . 替换成 [.]

func defangIPaddr(address string) string {
	// 直接无耻使用 strings.Replace() 方法
	return strings.Replace(address, ".", "[.]", -1)
}

func defangIPaddr2(address string) string {
	// 新建一个结构
	sb := strings.Builder{}
	for i := 0; i < len(address); i++ {
		c := address[i]
		if c != '.' {
			// 如果字符不是 . 就直接加入到 sb 中
			sb.WriteByte(c)
		} else {
			// 如果字符是 . 就把字符串 [.] 加入到 sb 中
			sb.Write([]byte("[.]"))
		}
	}
	return sb.String()
}

func main() {
	var address string = "1.1.1.1"
	fmt.Printf("%c\n", address[0])
	fmt.Println(defangIPaddr(address))
	fmt.Println(defangIPaddr2(address))
}
