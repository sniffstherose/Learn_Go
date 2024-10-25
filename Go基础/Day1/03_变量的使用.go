package main

import "fmt"

func main() {
	var a int
	fmt.Println("a =", a)
	a = 10
	fmt.Println("a =", a)

	var b int = 10
	b = 20
	fmt.Println("b =", b)

	c := 30
	fmt.Printf("c types is %T\n", c)
}