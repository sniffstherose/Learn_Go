package main

import "fmt"

/*
1. Go语言循环只支持for关键字，for有多种写法：
- 常规写法（和C语言一样，只不过少了括号）
- 无限循环：for后面无条件直接跟循环体
- 取代while：for后面只跟条件
- 配合range：for后面接range遍历数组
- 配合字符串：for后面接字符串，字符串底层为byte切片
- 剩下的写法还没学，
%todo

2. 结束循环的方式：
- 条件不满足
- break
- panic()
- goto
*/
func main() {
	/*
		这里仅演示特殊的情况：
		- 1.22 版本后的range变化。
	*/
	slice := []int{1, 2, 3, 4, 5}
	/*
		1.22 及以后的版本这里会随机顺序输出0~5，因为1.22以后的
		Goroutine会创建新的变量，不会共享变量
	*/
	for index, value := range slice {
		go func() {
			fmt.Printf("index: %d, value: %d\n", index, value)
		}()
	}
	// 支持整形循环
	for i := range 10 {
		fmt.Println(i)
	}
}
