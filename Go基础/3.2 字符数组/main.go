package main

import "fmt"

func main() {
	s1 := "localhost:8080"
	fmt.Println(s1)

	strbyte := []byte(s1)
	strbyte[len(s1)-1] = '1'
	fmt.Printf("%s\n", strbyte)

	s2 := string(strbyte)
	fmt.Println(s2)
}
