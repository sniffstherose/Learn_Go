package main

import "fmt"

func main() {
	// 变量声明使用var，常量声明使用const
	const a = 10	//常量自动类型推导用=不用:=
	const b int = 20
	fmt.Println("a =", a, "b =", b)

}