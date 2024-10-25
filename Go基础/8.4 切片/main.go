package main

import "fmt"

func main() {
	// var array [4]int
	// array[0]并未改变
	// change(array)
	// 为啥这样又能改变呢？因为[]留空就是声明一个切片，切片引用传递
	// 切片可以直接声明和赋值，也可以使用make关键字
	var array []int = []int{0, 0, 0}
	change(array)
	fmt.Println(array)

	a := [3]int{1, 2, 3}
	// a[1:2]生成了一个切片
	fmt.Println(a, a[0:2])

	// make一个切片
	var s []int = make([]int, 5, 10)
	fmt.Println(s)
}

func change(array []int) {
	array[0] = 10
}