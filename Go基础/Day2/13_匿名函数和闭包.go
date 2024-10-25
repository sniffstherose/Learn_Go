package main

import "fmt"

func main() {
	a := 10
	str := "nst"

	// 没写函数名的无参无返回值函数
	f1 := func() {
		fmt.Println(a)
		fmt.Println(str)
	}

	f1()

	// 匿名函数也可以变量赋值
	type funcType func()
	var f2 funcType
	f2 = f1
	f2()

	// 可以在定义处直接调用
	func(a, b int) {
		fmt.Printf("a = %d, b = %d\n", a, b)
	}(1, 2)	// 在这直接调用


	// 各种类型的函数都可以匿名
	x, y := func (i, j int) (max, min int) {
		if i > j {
			max = i
			min = j
		} else {
			max = j
			min = i
		}
		return
	} (10, 20)
	fmt.Printf("max = %d, min = %d\n", x, y)
}