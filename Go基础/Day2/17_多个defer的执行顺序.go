package main

import "fmt"

func main() {
	defer fmt.Println("aaaaa")
	defer fmt.Println("bbbbb")	// 后进先出
}