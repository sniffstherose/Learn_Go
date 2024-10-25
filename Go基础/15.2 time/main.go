package main

import (
	"fmt"
	"time"
)

// 有关时间的操作可以查阅资料
func main() {
	t := time.Now()
	fmt.Println(t.Format("2006-01-02 15:04:05"))
	milli := t.UnixMilli()
	second := t.Unix()
	fmt.Printf("秒值：%d, 毫秒值：%d\n", second, milli)
}
