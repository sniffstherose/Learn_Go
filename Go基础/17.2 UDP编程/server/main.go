package main

import (
	"fmt"
	"net"
)

/*
UDP比TCP简单，其步骤如下：
1. 监听地址
2. 直接接收数据
3. 回复数据
*/
func main() {
	// 第一步，监听
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 9999,
	})
	if err != nil {
		fmt.Printf("Listen failed, err:%v\n", err)
		return
	}

	// 第二步骤，循环读取消息
	for {
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Printf("Read failed from addr: %v, err: %v\n", addr, err)
			break
		}

		// 第三步，回复数据
		go func() {
			fmt.Printf("addr: %v data: %v count: %v\n", addr, string(data[:n]), n)
			_, err = listen.WriteToUDP([]byte("received success!"), addr)
			if err != nil {
				fmt.Printf("write filed, err: %v", err)
			}
		}()
	}

}
