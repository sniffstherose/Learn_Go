package main

import "fmt"

func myFunc01(a ...int) {	// 可变参数只能是函数的最后一个参数
	fmt.Println("len(a) =", len(a))
	for i, data := range a {
		fmt.Printf("a[%d] = %d\n", i, data)
	}
}

func main() {
	// 可变参数可以是0个,但如果有固定参数则一定要给固定参数传值，不能为空
	myFunc01()
	myFunc01(1, 2, 3)
}