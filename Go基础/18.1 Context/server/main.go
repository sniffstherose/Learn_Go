package main

import (
	"net/http"
	"time"
	"log"
)

func main() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("users"))
	})
	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(3 * time.Second)
		w.Write([]byte("products"))
	})
	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("orders"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
