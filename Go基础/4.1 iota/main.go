package main

import "fmt"

// iota为常量生成器，只能用于const
const (
	January = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

const (
	odd1 = 2*iota + 1
	odd2
	odd3
)

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday = iota
)

func main() {
	fmt.Println(December)
	fmt.Println(odd3)
	fmt.Println(Saturday)
}
