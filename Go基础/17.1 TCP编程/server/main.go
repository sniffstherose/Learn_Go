package main

import (
	"bufio"
	"fmt"
	"net"
)

/*
总体上看Go语言Tcp网络编程和C语言一样：
（0. 绑定地址）
1. 监听端口
2. 接收连接
3. 读写操作
4. 关闭连接
*/
func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("Read from client failed. Err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到来自Client的数据:", recvStr)
		conn.Write([]byte(recvStr))
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("Listen failed. Err:", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept failed. Err:", err)
			continue
		}
		go process(conn)
	}
}
