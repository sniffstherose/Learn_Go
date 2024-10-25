package main

import "fmt"

func add(a, b int) (result int) {
	result = a + b
	return
}

func minus(a, b int) (result int) {
	result = a - b
	return
}

// 函数也是类型，可以起别名
// 多态？
type funcType func(int, int) (int); 	// 函数类型，没有函数名，没有{}

func main() {
	var test funcType
	test = add
	result := test(10, 5)
	fmt.Println("result =", result)

	test = minus
	result = test(10, 5)
	fmt.Println("result =", result)
}