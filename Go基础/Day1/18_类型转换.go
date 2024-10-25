package main

import "fmt"

func main() {
	// 类型转换，有的类型不能相互转换，如bool和int这个和C++基本一样
	var ch byte
	ch = 'a'
	fmt.Printf("ch = %c\n", ch)
	var t int
	t = int(ch)
	fmt.Println("t =", t)
}