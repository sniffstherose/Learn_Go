package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// 读取文件
func readFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	buff := make([]byte, 1024)
	for {
		n, err := file.Read(buff)
		if err != nil {
			if err == io.EOF {
				fmt.Println("读完了")
				break
			}
			log.Fatal(err)
		}
		fmt.Printf("读取到的内容:%v\n", string(buff[:n]))
	}
}

func main() {
	// 读取字符串
	reader := strings.NewReader("你好啊，IO模块")
	nums := 0
	all := make([]byte, 0)
	for {
		p := make([]byte, 4)
		n, err := reader.Read(p)
		if err != nil {
			if err == io.EOF {
				fmt.Println("读取完成")
				log.Printf("读取的字节数：%d和内容：%v\n", nums, string(all[:nums]))
				break
			}
			log.Fatalf("读取错误:%v \n", err)
		}
		nums += n
		all = append(all, p...)
	}

	path := "Go基础\\16.1.1 IO包Reader\\main.go"
	readFile(path)
}
