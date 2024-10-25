package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Printf("请输入变量a和变量b: ")
	// fmt.Scanf("%d %d", &a, &b)
	// 还可以这样
	fmt.Scan(&a, &b)
	fmt.Println("a =", a, "b =", b)

}