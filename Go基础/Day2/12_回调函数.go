package main

import "fmt"

func add(a, b int) (result int) {
    result = a + b
    return
}

func minus(a, b int) (result int) {
    result = a - b
    return
}

func mul(a, b int) (result int) {
    result = a * b
    return
}

// 函数参数是函数类型，则这个函数就是回调函数
// 实现了多态，一个接口，多种形态
type funcType func(int, int) (int);
func cal(a, b int, test funcType) (result int) {
    fmt.Println("cal func")
    result = test(a, b) // 此时test是可以先暂时没有实现的
    return
}

func main() {
    a := cal(1, 1, minus)
    fmt.Println("a =", a)
}