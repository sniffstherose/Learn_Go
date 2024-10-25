package main

import "fmt"

func main() {
	// defer会延迟调用额，在main函数结束之前调用，所以现象会是先打印aaaaa
	defer fmt.Println("bbbbb")
	fmt.Println("aaaaa")
}