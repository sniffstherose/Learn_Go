package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func wr() {
	file, err := os.OpenFile("./test1.txt", os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	bufWriter := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		bufWriter.WriteString("hello\n")
	}
	bufWriter.Flush()
}

func rd() {
	file, err := os.OpenFile("./test1.txt", os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	bufReader := bufio.NewReader(file)
	for {
		line, _, err := bufReader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(line))
	}
}

func main() {
	wr()
	rd()
}
