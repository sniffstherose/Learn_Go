package main

import (
	"math/rand"
	"testing"
	"time"
)

func newSlice(n int) []int {
	rand.Seed(time.Now().UnixNano())
	// 注意，这里在生成切片的时候并没有指定容量
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}
func newSliceWithCap(n int) []int {
	rand.Seed(time.Now().UnixNano())
	// 注意，这里在生成切片的时候指定了容量
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

func BenchmarkNewSlice(b *testing.B) {
	for n := 0; n < b.N; n++ {
		newSlice(1000000)
	}
}

func BenchmarkNewSliceWithCap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		newSliceWithCap(1000000)
	}
}
