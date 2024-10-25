package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go performTask(ctx)
	time.Sleep(2 * time.Second)
	cancel() // 通知上面那个协程取消
	time.Sleep(time.Second)
}

func performTask(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Task cancelled.")
			return
		default:
			fmt.Println("Performing task...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
