package main

import "fmt"

/*
类型断言配合空接口使用，用于获取空接口变量的具体实例类型
语法：x.(T)，如v, t := x.(string)，表示断言x的类型是string，v为x的值，t为断言是否成功
还可以结合switch使用，具体细节看代码
*/
func main() {
	var x interface{} = "1234"
	v, ok := x.(string)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("类型断言失败")
	}
	switch v := x.(type) {
	case string:
		fmt.Println("x is a string:", v)
	case int:
		fmt.Println("x is a int:", v)
	default:
		fmt.Println("Unknown type")
	}
}
