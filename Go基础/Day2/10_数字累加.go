package main

import "fmt"

func test01() (sum int) {
	for i := 1; i <= 100; i++ {
		sum += i
	}
	return
}

func test02(a int) (sum int) {
	if a == 1 {
		return 1
	}
	return a + test02(a - 1)
}

func test03(a int) (sum int) {
	if a == 100 {
		return a		
	}
	return a + test03(a + 1)
}

func main() {
	var sum1 int
	sum1 = test01()
	var sum2 int
	sum2 = test02(100)
	var sum3 int
	sum3 = test03(1)
	fmt.Println("sum1 =", sum1)
	fmt.Println("sum2 =", sum2)
	fmt.Println("sum3 =", sum3)
}
