package main

import "fmt"

func main() {
	var a int = 10
	var b *int = &a

	*b = 20
	fmt.Println(a)

}