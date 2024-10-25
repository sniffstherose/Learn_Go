package main

import (
	"fmt"
	"sync"
	"time"
)

// 先不用纠结代码写法
func main() {
	start := time.Now().UnixMilli()
	task1()
	task2()
	end := time.Now().UnixMilli() - start
	fmt.Printf("串行运行总共用时:%d\n", end)

	start = time.Now().UnixMilli()
	var w sync.WaitGroup
	w.Add(2)
	go task3(&w)
	go task4(&w)
	w.Wait()
	end = time.Now().UnixMilli() - start
	fmt.Printf("并行运行总用时:%d", end)
}

func task1() {
	time.Sleep(time.Second * 3)
	fmt.Println("task1执行")
}

func task2() {
	time.Sleep(time.Second * 5)
	fmt.Println("task2执行")
}

func task3(w *sync.WaitGroup) {
	time.Sleep(time.Second * 3)
	fmt.Println("task3执行")
	w.Done()
}

func task4(w *sync.WaitGroup) {
	time.Sleep(time.Second * 5)
	fmt.Println("task4执行")
	w.Done()
}
