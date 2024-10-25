package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 9999,
	})
	if err != nil {
		fmt.Printf("failed dialudp, err: %v\n", err)
		return
	}

	for i := 0; i < 100; i++ {
		_, err = conn.Write([]byte("hello server"))
		if err != nil {
			fmt.Printf("Write failed, err: %v\n", err)
			return
		}

		result := make([]byte, 1024)
		n, remoteAddr, err := conn.ReadFromUDP(result)
		if err != nil {
			fmt.Printf("Read failed, err: %v\n", err)
			return
		}
		fmt.Printf("received from addr: %v data: %v\n", remoteAddr, string(result[:n]))
	}
}
