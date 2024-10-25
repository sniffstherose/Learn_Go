package main

import "fmt"

func main() {
	a := false
	b := true

	if a && b {
		fmt.Println("a && b == true")
	}
	if a || b {
		fmt.Println("a || b == true")
	}
	if !a {
		fmt.Println("!a == true")
	}
}
