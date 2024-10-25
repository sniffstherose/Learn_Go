package main

import "fmt"

func test01() {
	fmt.Println("test a =", a)
}

func main() {
	a = 10;
	fmt.Println("a =", a)
	test01()
}



// 定义在函数外面的变量，只需要保证在函数外面，前后没做硬性要求
// 不过为了代码可读性建议还是放在最上面
var a int