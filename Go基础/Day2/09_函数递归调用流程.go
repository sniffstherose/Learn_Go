package main

import "fmt"

func funca(a int) {
	if a == 1 {
		fmt.Println("a =", a)
	} else {
		funca(a - 1)
		fmt.Println("a =", a)
	}
}

func main() {
	funca(4)
	fmt.Println("main")
}
