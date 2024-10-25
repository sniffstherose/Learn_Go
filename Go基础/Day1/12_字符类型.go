package main

import "fmt"

func main () {
	var ch byte = 'c'
	fmt.Println("ch =", ch)	//输出对应的ascii码

	fmt.Printf("ch = %c, ch = %d\n", ch, ch)	//格式化输出

	ch1 := 'a'
	fmt.Printf("ch1 = %c\n", ch1)

	fmt.Printf("大写转小写:%c\n", 'A' + 32)
}