package main

import "fmt"

func main() {
	x := 10
	y := 10
	// 值传递不影响本身
	changeValue(x, y)
	fmt.Println(x, y)
	// 传递地址可以影响本身
	changePoint(&x, &y)
	fmt.Println(x, y)
	fmt.Println(yes(10, 20))
}

func changeValue(x, y int) {
	x = 50
	y = 100
}

func changePoint(x, y *int) {
	*x = 50
	*y = 100
}

// 函数返回值可以无，也可以有一个及多个，甚至可以命名，主要看看命名
func yes(x, y int) (a int, b int, info string) {
	a = 10
	b = 20
	info = "info"
	// return隐式返回
	return
}
