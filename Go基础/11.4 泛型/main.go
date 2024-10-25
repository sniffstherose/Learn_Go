package main

import "fmt"

/*
注意这里的T，这里T就是泛型的一个名字（想起别的名字一样可以）
T后面的int、float64...是约束，表示函数可以接受的类型，其实Go语言还有别的约束，比如any表示任意
comparable表示支持<和>操作符的类型，~int表示底层类型为int的类型等等。
泛型适用范围很广，还可以用在接口、方法、结构体中。
*/
func max[T int | float64 | float32 | int64](x, y T) T {
	if x > y {
		return x
	}
	return y
}

func main() {
	mInt := max[int]
	fmt.Println(mInt(2, 3))
	fmt.Println(max(3.14, 2.14))
}
