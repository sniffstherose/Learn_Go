package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	/*
		容量不够，会生成新的地址，由slice2存储
		注意切片扩容算法：
		- 容量小于256时两倍扩容
		- 大于256时按照一个公式扩容（倍率会越来越小）
		- 基于以上条件还会内存对齐
	*/
	slice2 := append(slice1, 6)
	fmt.Println(slice1, slice2)
}
