package main

import "fmt"

func main() {
	var (
		str1 string
		ch   byte
	)
	str1 = "Hello"
	ch = 'h'

	fmt.Println("ch =", ch, "str1 =", str1)

	fmt.Printf("str1[0] = %c, str1[1] = %c", str1[0], str1[1])

}
