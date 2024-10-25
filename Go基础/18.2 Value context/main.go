package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "userId", 123)
	go performTask(ctx)
	time.Sleep(time.Second)
}

func performTask(ctx context.Context) {
	userId := ctx.Value("userId")
	fmt.Println("userId:", userId)
}
