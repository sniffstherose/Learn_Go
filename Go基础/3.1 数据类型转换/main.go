package main

import (
	"fmt"
)

func main() {
	var a int = 10
	var b float64 = 3.14

	fmt.Println("a + b =", float64(a)+b)
	// go没有隐式转换，所有的类型转换需要显示声明
	c := 5.0
	d := int(3)
}
