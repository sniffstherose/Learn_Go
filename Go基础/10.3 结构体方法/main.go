package main

import "fmt"

func main() {
	user := User{}
	user.ModifyName("new")
	fmt.Println(user)
}

type User struct {
	Name string
}

/*
起的名字怪怪的，这里的(a *User)为接收器，其作用相当于接受传入类的对象
比如上面调用时的语法：user.ModifyName，此时接受到了user这个对象，并且
还是指针传递，可以修改对象本身的东西
（接收器可以是除了接口以外的任何类型）
方法和函数类似，无非就是有个所谓的接收器
*/
func (a *User) ModifyName(name string) {
	a.Name = name
}
