package main

import "fmt"

// 不过Goto个人用的少
func main() {
	if 20 > 10 {
		goto GotoTag1
	}
	return
GotoTag1:
	fmt.Println("Tag1")
	if 10 > 5 {
		goto GotoTag2
	}
	return
GotoTag2:
	fmt.Println("Tag2")
}
