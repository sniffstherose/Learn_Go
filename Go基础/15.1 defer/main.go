package main

import "fmt"

/*
defer是Go语言中的一个关键字，用于确保函数在执行结束时调用某个操作，通常用于资源清理、关闭文件、释放锁等。这是Go语言处理某些关键清理工作的强大工具。
其关键点为：
延迟执行：
defer语句会在其所在的函数返回之前执行。它们常用于保证某些操作（如关闭文件）在函数退出时进行。
先进后出（LIFO）：
多个defer语句的执行顺序是后进先出（LIFO）的。即最后一个defer语句最先执行。
函数参数的评估时间：
defer语句中的函数参数会在defer语句执行时评估，而不是在实际调用时。
*/
func main() {
	x := 10
	defer func() {
		x++
		fmt.Println("defer内后执行:", x)
	}()
	fmt.Println("defer外先执行:", x)
	return
}
