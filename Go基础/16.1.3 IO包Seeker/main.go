package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func writeFile(content string) {
	file, err := os.OpenFile("./test.txt", os.O_RDWR, 655)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	_, err = file.Seek(4, io.SeekStart)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = file.Write([]byte(content))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var temp string = "nihao"
	writeFile(temp)
}
