package main
import "fmt"
func main() {
	str := "abc"
	// 通过传统for循环
	for i := 0; i < len(str); i++ {
		fmt.Printf("str[%d] = %c\n", i, str[i])
	}
	// 通过range
	for i, data := range str {
		fmt.Printf("str[%d] = %c\n", i, data)
	}
	// 可以丢弃一些参数,用"_",如果没有加"_"默认也是丢弃
	for i, _ := range str {
		fmt.Printf("str[%d] = %c\n", i, str[i])
	}
}