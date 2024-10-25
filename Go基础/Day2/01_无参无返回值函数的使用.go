package main

import "fmt"

// go语言不需要再起始地声明函数
func main() {
	myFunc()
}

func myFunc() {
	a := 555
	fmt.Println("a =", a)
}