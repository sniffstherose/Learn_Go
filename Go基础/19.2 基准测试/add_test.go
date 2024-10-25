package main

import (
	"fmt"
	"strconv"
	"testing"
)

// 以Benchmark开头的函数
func BenchmarkSprintf(b *testing.B) {
	num := 10

	b.ResetTimer() // 重置计时器开始测试
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", num)
	}
}

func BenchmarkFormat(b *testing.B) {
	num := int64(10)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strconv.FormatInt(num, 10)
	}
}
