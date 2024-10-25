package main

import "fmt"

func main() {
	// 匿名函数
	funA := func() int{
		return 20
	}
	// funA其实就是该函数的别名了
	fmt.Println(funA());
	// 匿名函数可以直接调用，只需在最后加个括号
	func() {
		fmt.Println("这是一个匿名函数")
	} ()
}