package main

import (
	"fmt"
	"time"
)

// 有关进程、线程、协程的内容网上有很多博客
func main() {
	for i := 0; i < 10; i++ {
		go hello(i)
	}
	fmt.Println("main goroutine done!")
	time.Sleep(time.Second * 2) // 如果不加休眠部分的代码，在主协程结束后便会退出整个进程，所以得加
}

func hello(i int) {
	fmt.Println("Hello Goroutine!", i)
}
