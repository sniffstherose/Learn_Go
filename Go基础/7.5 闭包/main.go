package main

import "fmt"

func main() {
	generator := playerGen("张三")
	name, hp := generator()
	fmt.Println(name, hp)
}

// 闭包可以简单理解为函数内部的匿名函数
func playerGen(name string) func() (string, int) {
	hp := 150
	// 返回一个闭包
	return func() (string, int) {
		return name, hp
	}
}