package main

import (
	"fmt"
	"os"
)

func writeFile(content string) {
	// 新建文件
	file, err := os.Create("./test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	_, err = file.Write([]byte(content))
	if err != nil {
		return
	}

}

func main() {
	var content string = "Hello IO!!!"
	writeFile(content)
}
