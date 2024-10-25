package main

import "fmt"
import "os"	// 需要导入头文件

func main() {
	list := os.Args

	n := len(list)
	fmt.Println("n =", n)

	for i, data := range list {
		fmt.Printf("list[%d] = %s\n", i, data)
	}
}