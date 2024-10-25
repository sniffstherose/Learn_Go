package main

import "fmt"

func main() {
	fmt.Println(max(10, 20))
	fmt.Println(max(-1, -2))
	a := 10
	b := 20
	a, b = swap(a, b)
	fmt.Printf("%d, %d\n", a, b)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func swap(a, b int) (int, int) {
	return b, a
}