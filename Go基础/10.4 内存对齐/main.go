package main

import (
	"fmt"
	"unsafe"
)

type test1 struct {
	a bool   // 1
	b int32  // 4
	c string // 16
}

type test2 struct {
	a int32
	b string
	c bool
}

func main() {
	var t1 test1
	var t2 test2

	fmt.Println("t1 size is ", unsafe.Sizeof(t1)) // 24
	fmt.Println("t2 size is ", unsafe.Sizeof(t2)) // 32
}
