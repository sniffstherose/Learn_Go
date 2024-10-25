package main

import "fmt"

func myFunc01() int {
	return 666
}

// 最常用写法
func myFunc02() (result int) {
	result = 555
	return // 有返回值，如果给返回值赋值则不用跟东西
}

func main() {
	var a int
	a = myFunc01()
	fmt.Println("a =", a)

	// 自动类型推导
	b := myFunc01()
	fmt.Println("b =", b)

	c := myFunc02()
	fmt.Println("c =", c)
}
