package main

import "fmt"

/*
	copy( destSlice, srcSlice []T) int )
	- 作用：用于将srcSlice切片复制到destSlice上
	- 属于深拷贝（值拷贝，不会产生地址冲突）
	- 按照容量小的容量复制
*/
func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	// 只会复制前3个元素
	copy(slice2, slice1)
	// 修改slice2不会影响slice1
	slice2[0] = 0
	fmt.Println(slice1, slice2)
}