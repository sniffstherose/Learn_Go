package main

import "fmt"

func main () {
	var a bool
	
	fmt.Println("a =", a)	//bool的零值为false

	a = true

	fmt.Println("a =", a)

	// 自动类型推导
	var b = true
	fmt.Println("b =", b)
	
	c := true
	fmt.Println("c =", c)
}