package main

import "fmt"

func main () {
	var (
		a int
		b float64
	)

	a = 10
	b = 3.14

	fmt.Println("a =", a)
	fmt.Println("b =", b)

	const (
		c  = 10
		d  = 1.1
	)
	// 也可以自动类型推导

	fmt.Println("c =", c)
	fmt.Println("d =", d)
}