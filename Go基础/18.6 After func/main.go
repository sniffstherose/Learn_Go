package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // 确保在函数结束时取消context
	// 设置一个2秒后执行的函数
	stop := context.AfterFunc(ctx, func() {
		fmt.Println("AfterFunc executed")
	})
	go performTask(ctx, stop)
	// 阻塞主goroutine，防止程序立即退出
	time.Sleep(3 * time.Second)
}

func performTask(ctx context.Context, stop func() bool) {
	select {
	case <-ctx.Done():
		fmt.Println("Context canceled:", ctx.Err())
		fmt.Println("stop:", stop())
	}
}
