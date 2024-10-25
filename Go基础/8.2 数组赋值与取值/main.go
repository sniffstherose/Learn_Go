package main

import "fmt"

/*
赋值：
	初始化方式：
	1. 可以赋全值
	2. 可以只赋从下标0开始连续几个值
	3. 可以自动类型推导
	4. 可以不写具体长度
	5. 可以给指定索引赋值
	索引方式：
	直接将下标对应的元素赋值
取值：
	索引方式：
	直接将下标对应元素取出来
	for range：

*/

func main() {
	// 初始化方式赋值
	var arr1 [3]int = [3]int{1, 2, 3}
	var arr2 [3]int = [3]int{1, 2}
	arr3 := [3]int{1, 2, 3}
	arr4 := [...]int{1, 2, 3, 4}
	arr5 := [3]int{2: 3}
	fmt.Println(arr1, arr2, arr3, arr4, arr5)
	// 索引方式赋值
	arr1[0] = 0
	arr1[1] = 0
	arr1[2] = 0
	fmt.Println(arr1)

	// 索引方式取值
	fmt.Println(arr4[3])
	// for range方式取值
	for index, value := range arr4 {
		fmt.Printf("索引：%d， 值：%d\n", index, value)
	}
}