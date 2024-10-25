package main

import "fmt"

func main() {
	var a int = 0
	fmt.Println("请输入a的值：")
	fmt.Scan(&a)
	switch {
	case a >= 90:
		fmt.Println("优秀")
	case a >= 70:
		fmt.Println("良好")
	case a >= 60:
		fmt.Println("及格")
	case a < 60:
		fmt.Println("不及格")
		fallthrough
	default:
		fmt.Println("default")
	}
}
