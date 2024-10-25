package main

import "fmt"

func main () {
	var f1 float32
	f1 = 3.14
	fmt.Println("f1 =", f1)

	// 自动类型推导, 其类型为float64
	f2 := 3.14
	fmt.Printf("f2 type is %T\n", f2)
}