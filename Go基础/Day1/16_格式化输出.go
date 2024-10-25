package main

import "fmt"

func main() {
	a := 10
	b := "abc"
	c := 'a'
	d := 1.1

	fmt.Printf("%T, %T, %T, %T\n", a, b, c, d)
	fmt.Printf("%d, %s, %c, %f\n", a, b, c, d)
	// %v自动格式化输出，不过可能不是很智能，比如%c会输出成整数
	fmt.Printf("%v, %v, %v, %v", a, b, c, d)
}