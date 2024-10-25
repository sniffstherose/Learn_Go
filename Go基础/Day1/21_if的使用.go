package main

import "fmt"

func main() {
	s := "屌丝"
	if s == "王思聪" {
		fmt.Println("左手一个妹子，右手一个大妈")
	}

	// if支持一个初始化语句，初始化语句和条件用;分隔
	if a := 10; a == 10 {
		fmt.Println("a == 10")
	}
}