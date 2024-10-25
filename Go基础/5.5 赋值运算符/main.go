package main

import "fmt"

func main() {
	var a int = 10
	fmt.Println("a =", a)
	a += 10
	fmt.Println("a =", a)
	a -= 10
	fmt.Println("a =", a)
	a *= 10
	fmt.Println("a =", a)
	a /= 10
	fmt.Println("a =", a)
	a %= 3
	fmt.Println("a =", a)
	a <<= 1
	fmt.Println("a =", a)
	a >>= 1
	fmt.Println("a =", a)
	a &= 1
	fmt.Println("a =", a)
	a |= 1
	fmt.Println("a =", a)
	a ^= 1
	fmt.Println("a =", a)
}