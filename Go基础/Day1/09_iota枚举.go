package main

import "fmt"

func main() {
	// 1.iota
	const (
		a = iota	//0
		b = iota	//1
		c = iota	//2
	)

	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)

	// 2.只写一个iota
	const (
		a1 = iota
		b1 = iota
		c1 = iota
	)

	fmt.Printf("a1 = %d, b1 = %d, c1 = %d\n", a1, b1, c1)

	// 3.iota遇到const重置为0
	const d = iota
	fmt.Println("d =", d)

	// 4.同一行的值都一样
	const (
		i = iota
		j1, j2, j3 = iota, iota, iota
		k = iota
	)

	fmt.Printf("i = %d, j1 = %d, j2 = %d, j3 = %d, k = %d", i, j1, j2, j3, k)
}