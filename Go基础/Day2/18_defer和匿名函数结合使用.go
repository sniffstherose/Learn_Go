package main

import "fmt"

func main01() {
	a := 10
	b := 20

	defer func() {
		fmt.Printf("a = %d, b = %d\n", a, b)
	} ()
	a = 100
	b = 200
	fmt.Printf("a = %d, b = %d\n", a, b)
}

func main() {
	a := 10
	b := 20

	// 会在程序结束前再执行
	defer func(a, b int) {
		fmt.Printf("a = %d, b = %d\n", a, b)
	} (a, b)	// 调用，这里先传参，最后再调用
	a = 100
	b = 200
	fmt.Printf("a = %d, b = %d\n", a, b)
}