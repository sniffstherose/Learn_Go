package main

import "fmt"

var a int

func main() {
	var a byte	// 局部变量
	fmt.Printf("%T\n", a)	// uint8，就近原则

	{
		var a float32
		fmt.Printf("%T\n", a)	// float32
	}

	test()	// int，因为作用域的原因，会使用全局变量a
}

func test() {
	fmt.Printf("%T\n", a)	// int
}