package main
import "fmt"
func main() {
	// 可以使用type来为类型起一个别名
	type bigint int64
	
	var a bigint
	fmt.Printf("a types is %T\n", a)

	// 也可以同时定义多个type
	type (
		long int64
		char byte
	)

	var b long = 11
	var c char = 'a'
	fmt.Printf("b = %d, c = %c\n", b, c)
}