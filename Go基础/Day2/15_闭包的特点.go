package main

import "fmt"

func test01() (int){
	var x int	// 函数被调用时x才分配空间并初始化为0，函数结束x就会被销毁
	x++
	return x * x	// 用完即释放
}

// 返回值是一个函数类型
func test02() func() (int) {
	var x int

	return func() (int) {
		x++
		return x * 
	}
}

func main() {
	a := test01()
	fmt.Println("a =", a)	// 1
	a = test01()
	fmt.Println("a =", a)	// 1

	f := test02()
	// 闭包使用变量时不关心作用域，只要闭包还在使用则该变量一直都在
	fmt.Println(f())	// 1
	fmt.Println(f())	// 4
	fmt.Println(f())	// 9
	fmt.Println(f())	// 16
}