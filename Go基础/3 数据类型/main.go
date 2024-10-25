package main

import "fmt"

func main() {
	// 整型
	var health int64
	var level int16
	fmt.Printf("玩家血量：%d, 等级：%d", health, level)
	// 浮点型
	floatstr1 := 3.14
	fmt.Printf("%.2f\n", floatstr1)
	var f float32 = 1 << 24
	fmt.Println(f == f+1)
	var a = .12345
	var b = 1.
	fmt.Printf("%.5f, %.1f\n", a, b)
	var avogadro = 6.02214129e23 // 阿伏伽德罗常数
	var planck = 6.62606957e-34  // 普朗克常数
	fmt.Printf("%f\n,%.35f\n", avogadro, planck)

	// 字符型
	var c byte = 'A'
	var d byte = 65
	var e byte = '\x41'
	var g byte = '\101'
	fmt.Printf("%c, %c, %c, %c\n", c, d, e, g)

	// Unicode
	var ni rune = '你'
	var hao rune = '\u597D'
	var hao1 rune = '好'
	fmt.Printf("%c, %c, %U\n", ni, hao, hao1)

	// 字符串类型
	var str1 string = "你好\nGolang"
	var str2 string = `你好\nGolang`
	fmt.Println(str1)
	// 使用``包裹则会直接原样输出
	fmt.Println(str2)
	str3 := "abcde"
	// str3[0] = "b", 不行，只能重新赋值
	str3 = "defgh"
	fmt.Println(str3)
	str4 := "hello,码神之路的go教程" // 3 * 7 + 8 = 29
	fmt.Printf("str4.length: %d\n", len(str4))

	// 类型转换
	var i1 int = 10
	var f1 float64 = 3.14
	//fmt.Println("i1 + f1 =", i1 + f1)	// 错误
	fmt.Println("i1 + f1 =", float64(i1)+f1)
	
}
