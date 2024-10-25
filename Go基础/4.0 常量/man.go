package main

import "fmt"

// 常量只能定义布尔型，整型，浮点型和底层类型为这些类型的（比如byte，rune）
const PI float64 = 3.1415926

func main() {
	fmt.Println(PI)
}