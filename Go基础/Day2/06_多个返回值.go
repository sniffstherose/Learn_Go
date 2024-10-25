package main

import "fmt"

func myFunc01() (int, int, int) {
	return 1, 2, 3
}

// go风格推荐写法
func myFunc02() (a int, b int, c int) {
	a, b, c = 111, 222, 333
	return
}

func myFunc03() (a, b, c int) {
	a, b, c = 111, 222, 333
	return
}


func main() {
	a, b, c := myFunc01()
	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)
	
	d, e, f := myFunc03()
	fmt.Printf("d = %d, e = %d, f = %d\n", d, e, f)
}
