package main
import "fmt"
func main() {
	// 求1 ~ 100累加和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i;
	}
	fmt.Println("sum =", sum)
}