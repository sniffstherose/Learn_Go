package main

import "fmt"

// 没有方法的接口称之为空接口，任何类型都实现了空接口，所以可以用任何类型的实例赋值给空接口类型的变量
type Any interface{}

// Go语言默认有空接口类型别名：any
func main() {
	var a Any
	a = 10
	var b any = 11.2
	fmt.Println(a)
	fmt.Println(b)
}
