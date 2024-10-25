package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x      int64
	wg     sync.WaitGroup
	rwlock sync.RWMutex
)

// 修改临界资源，如果不加锁，两个协程最后的结果并不为10000
// 读多写少的场景应当使用读写锁提高运行效率
func write() {
	// lock.Lock() 加互斥锁
	rwlock.Lock() // 加写锁
	x += 1
	time.Sleep(10 * time.Millisecond)
	rwlock.Unlock() // 解写锁
	// lock.Unlock()   解互斥锁
	wg.Done()
}
func read() {
	// lock.Lock() 加互斥锁
	rwlock.RLock() // 加读锁
	time.Sleep(time.Millisecond)
	rwlock.RUnlock() // 解读锁
	// lock.Unlock()   解互斥锁
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()
	fmt.Println(x)
	end := time.Now()
	fmt.Println("总用时:", end.Sub(start))
}
