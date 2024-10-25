package main

import "fmt"

func main() {
	// 外部变量
	a := 10
	str := "mike"

	// 闭包
	func() {
		// 闭包是以引用传递外部变量
		a = 666
		str = "go"
		fmt.Printf("闭包内部 a = %d, str = %s\n", a, str)
	}()	// 不要忘了调用
	fmt.Printf("闭包外部 a = %d, str = %s\n", a, str)
}