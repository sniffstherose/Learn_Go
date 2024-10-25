package main

import "fmt"

func main() {
	// go语言中函数可以当做一种数据类型
	var swap func(x, y int) (int, int)
	swap = func(x, y int) (int, int) {
		return y, x
	}
	fmt.Println(swap(10, 20))

	f := call(swap)
	f(10)

}

// 函数既然是变量，则可以当做参数传递，也可作为返回值
func call(swap func(x, y int) (int, int)) func(x int) {
	y, _ := swap(10, 20)
	return func(x int) {
		fmt.Println(x + y)
	}
}
