package main

import "fmt"

func myFunc01(args ...int) {
	for _, data := range args {
		fmt.Println("data =", data)
	}
}

func test(args ...int) {
	// 把全部元素传递给myFunc01
	myFunc01(args...)
	// 把前两个元素传递给myFunc01
	fmt.Println("********************************")
	myFunc01(args[:2]...)	// 不包括args[2]
	// 把后两个元素传递给MyFunc01
	fmt.Println("********************************")
	myFunc01(args[2:]...)	// 包括args[2]
}

func main() {
	test(1, 2, 3, 4)
}
