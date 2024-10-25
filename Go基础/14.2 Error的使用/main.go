package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	Msg string
}

func (e *MyError) Error() string {
	return e.Msg
}

var MoneyNotEnough error = errors.New("余额不足")
var MoneyNot = &MyError{Msg: "余额不足"}

func pay(money float64) error {
	if money < 10 {
		return MoneyNot // 注意这里返回的是同一个实例.
	}
	return nil
}

/*
As的好处:
- 类型断言和转换：errors.As 能够将错误链中的错误转换为特定的错误类型。这在处理多个自定义错误类型时非常有用，简化了类型断言的过程。
- 处理错误链：errors.As 会检查整个错误链（即嵌套的、包装的错误），直到找到匹配的类型。这让你可以对被包装的错误进行处理，而不用关心错误在链中的深度。
- 增强可读性和维护性：使用errors.As 可以使代码更加简洁、清晰，不需要手动编写复杂的类型断言逻辑。
- 兼容性：通过提供一致的方式处理错误，errors.As 使得不同模块和库之间的错误处理更加一致和兼容
*/
func main() {
	err := pay(5)
	var myError *MyError
	if errors.As(err, &myError) {
		fmt.Println("转换成功")
		fmt.Println(myError.Error())
	} else {
		fmt.Println("不是此error类型")
	}

	// fmt.Errorf()对错误进行封装
	newErr := fmt.Errorf("error: %w", err)
	if errors.Is(newErr, err) {
		fmt.Println("封装后的错误与原错误相等")
	}

	// errors.Join进行多个错误封装
	err1 := errors.Join(err, newErr)
	if errors.Is(err1, err) {
		fmt.Println("err1 is err")
	}
	if errors.Is(err1, newErr) {
		fmt.Println("err is newErr")
	}
}
