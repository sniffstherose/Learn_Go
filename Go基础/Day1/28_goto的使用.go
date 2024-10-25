package main

import "fmt"

func main() {
	fmt.Println("11111111111111")
	goto End	// 直接跳转到标签处执行

	fmt.Println("22222222222222")

End:
	fmt.Println("33333333333333")
}