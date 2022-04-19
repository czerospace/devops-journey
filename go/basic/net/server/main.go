package main

import (
	"fmt"
	"net"
)

// 网路服务器端

func process(conn net.Conn) {
	// 连接完一定要关闭:
	defer conn.Close()
	for {
		// 创建一个切片，准备:将读取的数据放入切片
		buf := make([]byte, 1024)
		// 从 conn 连接中读取数据:
		n, err := conn.Read(buf)
		if err != nil {
			return
		}
		// 将读取内容在服务器端输出:
		// buf 长度为 1024,只输出 0-n的内容
		fmt.Println(string(buf[0:n]))
	}
}

func main() {
	fmt.Println("server starting...")
	// 进行监听: 需要指定服务器端 TCP 协议，服务器端的 IP+PORT
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("Listening faild: ", err)
		return
	}

	// 监听成功以后
	// 循环等待客户端的链接
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("connecting faild: ", err)
		} else {
			// 连接成功
			fmt.Printf("connecting succeed,con=%v,client info: %v\n", conn, conn.RemoteAddr().String())
		}

		// 准备一个协程，协程处理客户端服务请求:
		// 不同客户端的请求，连接 conn 不一样的
		go process(conn)
	}

}
