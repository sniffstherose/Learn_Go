package main

import "fmt"

func main() {
	// 此处的int为key的类型，string为value的类型
	var m map[int]string = map[int]string{}
	/*
		也可以使用make方式：
		var m = make( map[int]string, 2 )
		- 这里的2是容量
		- map是引用类型，零值3为nil，所以需要初始化了使用
	*/
	var n = make(map[int]string, 3)
	m[1] = "1"
	m[2] = "2"
	n[1] = "1"
	n[3] = "3"
	n[4] = "4"
	// 引用
	change(m)
	fmt.Println(n, m)

	/*
		map有两种取值方式：
		1. 通过key取值：map[key]，其表示的就是map中key对应的value的值
		2. 通过range取值：使用范围循环range，但注意返回的结果并不是有序的
	*/
	m1 := make(map[string]string)
	m1["key"] = "value"
	m1["key1"] = "value1"
	// 这里的ok可以用于判断key是否存在m1中
	v, ok := m1["key"]
	if ok {
		fmt.Println(v)
	}
	// 通过range取值
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	// 删除键值对用delete(map, key)
	delete(m1, "key1")
	fmt.Println(m1)

	// map中可以实现一对多的关系，需要将value设置为切片类型
	m2 := make(map[string][]int)
	v1 := []int{
		1, 2, 3, 4, 5,
	}
	m2["key1"] = v1
	fmt.Println(m2)
}

func change(m map[int]string) {
	m[2] = "改变了"
}
