package main

import "fmt"

type User struct {
	Name    string
	Age     int
	Address Address
}

type Address struct {
	City string
}

type Human struct {
	string
}

func main() {
	/*
		听过关键字type定义结构体，结构体是复合类型，具体语法为：
		type 结构体名 struct {
			成员 成员类型
			...
		}
		注意此处的结构体名在同一个包内不能重复
		成员名首字母大写代表public（所有包可见），否则private（本包可见）
		结构体之间可以嵌套
		结构体中的字段可以没有名字，称之为匿名字段，如果要用到直接输入其类型即可
		注意，由于匿名字段以类型区分，所以不可以有同类型的匿名字段
	*/

	// 结构体的使用，可以通过 = 赋值，通过.取值
	// 这里也可以声明一个结构体指针：user := &User{}
	user := User{}
	user.Name = "张三"
	user.Age = 18
	user.Address.City = "北京"
	fmt.Printf("Name：%s, Age：%d, Addr：%s\n", user.Name, user.Age, user.Address.City)
	/*
		和其他语言一样，Go语言传参最好也使用指针
		可以避免不必要的内存开销，如果是普通结构
		体可以使用*T，如果是切片可以使用[]*T
	*/
	change(&user)
	fmt.Println(user.Name, user.Age)
}

func change(user *User) {
	user.Name = "change"
	user.Age = 99
}
