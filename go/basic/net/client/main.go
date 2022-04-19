package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// 网络客户端

func main() {
	fmt.Println("client starting...")
	// 调用 Dial 函数: 参数需要指定 tcp 协议，需要指定服务器端的 IP+PORT
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("connect failed...")
		return
	}
	fmt.Println("connecting succeed：", conn)

	// 通过客户端发送单行数据，然后退出:
	// os.Stdin 代表终端标准输入
	reader := bufio.NewReader(os.Stdin)

	// 从终端读取一行用户输入的信息:
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("client input failed err: ", err)
	}

	// 将 str 数据发送给服务器:
	n, err := conn.Write([]byte(str))
	if err != nil {
		fmt.Println("connecting faild err: ", err)
	}
	fmt.Printf("send data to server succeed,total of data is %d", n)
}
