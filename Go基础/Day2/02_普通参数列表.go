package main

import "fmt"

func myFunc01(a int, b int) {	// 还可以写成func myFunc01(a, b int) {}
	fmt.Printf("a = %d, b = %d\n", a, b)
}

func main() {
	myFunc01(555, 666)
}