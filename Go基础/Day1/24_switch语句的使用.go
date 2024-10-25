package main

import "fmt"

func main() {
	var a int
	fmt.Println("请输入a的值")
	fmt.Scan(&a)

	switch a { // switch可以没有条件，case里放<,>等，也可以switch里放一个初始化
	case 1: // 可以这样写多条case 1,2,3:
		fmt.Println("a = 1")
		// go语言默认会有break，所以可以不加
	case 2:
		fmt.Println("a = 2")
	case 3:
		fmt.Println("a = 3")
		fallthrough // 加上这个后这条case和c中没加break效果一样
	case 4:
		fmt.Println("a = 4")
	default:
		fmt.Println("其他")
	}
}
