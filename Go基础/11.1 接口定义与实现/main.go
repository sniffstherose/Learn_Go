package main

import "fmt"

/*
定义接口用interface关键字，和java一样
接口是一系列方法的集合，不需要在接口里具体实现方法
*/
type Animal interface {
	Say()
	Move()
	Jump()
}

/*
Dog结构体实现的方法和接口的一致，所以Dog实现了接口
Go语言中不用显示的写出implements，编译器会自动判断。
*/
type Dog struct {
}

// 
func (d *Dog) Say() {
	fmt.Println("汪汪汪")
}
func (d *Dog) Move() {
	fmt.Println("狗在跑")
}
func (d *Dog) Jump() {
	fmt.Println("狗在跳")
}

type Cat struct {
}

func (c *Cat) Say() {
	fmt.Println("喵喵喵")
}

func main() {
	// 这里的声明语句也体现了Go语言的多态，可以用实现了接口的类型的实例来初始化接口类型的变量
	var dog Animal = &Dog{}
	// 这里的结果为nil，接口的0值也是nil，想要使用必须初始化
	dog.Say()
}
