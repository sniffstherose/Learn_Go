package main
import "fmt"
func main() {
	a := 10


	// 这样写会更好，应为条件判断次数会更少，下面那样每个条件都判断了一次
	if a == 10 {
		fmt.Println("a == 10")
	} else if a < 10 {
		fmt.Println("a < 10")
	} else if a > 10 {
		fmt.Println("a > 10")
	} else {
		fmt.Println("不可能")
	}

	b := 10
	if b == 10 {
		fmt.Println("b == 10")
	}
	if b > 10 {
		fmt.Println("b > 10")
	}
	if b < 10 {
		fmt.Println("b < 10")
	}
}