package main

import "fmt"

func main() {
	var a *int
	old := 10
	a = &old

	fmt.Printf("打印a的值：%v\n", a)
	fmt.Println("打印a指向地址的值：", *a)
	*a = 20
	fmt.Println("打印old的值：", old)

	// 使用new分配内存
	var b *int = new(int)
	*b = 10
	fmt.Println("打印b指向地址的值：", *b)
}
