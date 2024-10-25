package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// 客户端相对简单，只需要调用Dial函数连接即可

func main() {
	// 1. 连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("err :", err)
		return
	}
	defer conn.Close() // 确保在函数结束时关闭连接

	// 2. 创建一个从标准输入读取数据的读取器
	inputReader := bufio.NewReader(os.Stdin)

	// 3. 进入一个无限循环，持续读取用户输入并发送到服务器
	for {
		// 读取用户输入的一行数据
		input, _ := inputReader.ReadString('\n')
		// 去除输入数据末尾的回车换行符
		inputInfo := strings.Trim(input, "\r\n")
		// 如果输入的是"Q"或"q"，则退出循环
		if strings.ToUpper(inputInfo) == "Q" {
			return
		}
		// 将用户输入的数据发送到服务器
		_, err = conn.Write([]byte(inputInfo))
		if err != nil {
			return
		}
		// 创建一个512字节的缓冲区，用于接收服务器响应的数据
		buf := [512]byte{}
		// 读取服务器响应的数据到缓冲区中
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("recv failed, err:", err)
			return
		}
		// 打印服务器的响应数据
		fmt.Println(string(buf[:n]))
	}
}
