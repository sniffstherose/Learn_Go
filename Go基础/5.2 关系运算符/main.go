package main

import "fmt"

func main() {
	a := 10
	b := 21
	compare(a, b)
}

func compare(a, b int) {
	if a == b {
		fmt.Println("a等于b")
	}
	if a != b {
		fmt.Println("a不等于b")
	}
	if a > b {
		fmt.Println("a大于b")
	}
	if a < b {
		fmt.Println("a小于b")
	}
	if a >= b {
		fmt.Println("a大于等于b")
	}
	if a <= b {
		fmt.Println("a小于等于b")
	}
}
