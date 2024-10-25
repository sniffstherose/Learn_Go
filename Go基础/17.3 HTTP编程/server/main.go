package main

import (
	"fmt"
	"net/http"
)

// handler函数
func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr, "连接成功")
	fmt.Println("method:", r.Method)
	fmt.Println("url:", r.URL.Path)
	fmt.Println("header:", r.Header)
	fmt.Println("body:", r.Body)

	w.Write([]byte("你好，HTTP客户端"))

}

func main() {
	// 注册处理函数
	http.HandleFunc("/go", myHandler)
	// 启动服务器（默认服务器）
	http.ListenAndServe("127.0.0.1:8080", nil)
}
