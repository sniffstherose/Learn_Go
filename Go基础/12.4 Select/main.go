package main

import (
	"fmt"
	"time"
)

/*
select在Go语言中是一种多路复用机制
使用select case语句可以在多个通道上等待并处理消息
某个通道有消息则对应case语句会执行，若同时有多个
case可执行的话则随机执行一个
*/
func main() {
	ch := make(chan int)
	go func() {
		time.Sleep(time.Second * 3)
		ch <- 1
	}()

	select {
	case data, ok := <-ch:
		if ok {
			fmt.Println("接收到数据:", data)
		} else {
			fmt.Println("通道已被关闭")
		}
	case <-time.After(time.Second * 2):
		fmt.Println("超时！")
	}
}
