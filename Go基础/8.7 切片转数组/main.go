package main

import "fmt"

func main() {
	slice := make([]int, 2, 4)
	// go 1.17版本特性
	arr := *(*[2]int)(slice)
	fmt.Println(arr)
	// go 1.20新特性
	arr1 := [2]int(slice)
	fmt.Println(arr1)
}