package main

import (
	"errors"
	"log"
)

var MoneyNotEnough error = errors.New("money not enough")

func pay(money float64) error {
	if money < 10 {
		return MoneyNotEnough
	}
	return nil
}

/*
在Go语言中，panic是一种用于处理严重错误的机制，它会导致程序的运行中止，并开始执行所有延迟（defer）的函数。以下是panic的关键点：
使用场景
通常用于处理程序无法恢复的严重错误，例如数组越界、空指针引用等。
不建议在常规错误处理中使用panic，而是应优先使用返回错误值的方式。
*/
func main() {
	err := pay(5)
	// 可以使用recover函数捕获panic, recover只有在defer函数中有效
	defer func() {
		if r := recover(); r != nil {
			log.Printf("panic:%+v\n", r)
		}
	}()
	if err != nil {
		panic("余额不足")
	}
}
