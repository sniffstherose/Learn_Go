package main

import "fmt"

func funca (a int) {
	funcb(a - 1)
	fmt.Println("a =", a)
}

func funcb (b int) {
	funcc(b - 1)
	fmt.Println("b =", b)
}

func funcc (c int) {
	fmt.Println("c =", c)
}

func main() {
	funca(3)
	fmt.Println("main")
}
