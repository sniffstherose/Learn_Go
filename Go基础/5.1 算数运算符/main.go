package main

import "fmt"

func main() {
	a := 21
	b := 10
	c := 0
	c = a + b
	fmt.Println(c)
	c = a - b
	fmt.Println(c)
	c = a * b
	fmt.Println(c)
	div := float64(a) / float64(b)
	fmt.Println(div)
	c = a % b
	fmt.Println(c)
}
