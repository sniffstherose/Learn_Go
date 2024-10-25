package main

import "fmt"

func test1() {
	a := 111
	fmt.Println("a =", a)
}

func main() {
	// 和C/C++类似，大括号就是变量的作用域，离开变量的作用域自动释放内存
	// a = 10 undefined
	if flag := 3; flag == 3 {
		fmt.Println("flag =", flag)
	}
	// flag = 4 undefined
}