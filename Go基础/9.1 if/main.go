package main

import "fmt"

func main() {
	a := 55
	if a > 100 {
		fmt.Println("a > 100")
	} else if a > 50 {
		fmt.Println("a > 50")
	} else if a < 20 {
		fmt.Println("a < 20")
	} else {
		fmt.Println("a")
	}

	// 可以这样写：if a := 100; a > 100 {}，这样将a的作用域限制在if内部，可以增强程序稳定性
}
