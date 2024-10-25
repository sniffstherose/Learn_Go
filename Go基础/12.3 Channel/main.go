package main

import "fmt"

/*
channel（通道）可用于协程间通信：
- 语法格式：var 变量名 chan 类型
- chan是应用类型，零值为nil，所以需要初始化才能使用
- 初始化语法：make(chan 类型, [缓冲大小])
- 通道有发送、接收、关闭三种操作，发送都用<-符号，非常的形象
- 关闭通道不一定必须，可以被垃圾回收机制回收
通道有两种：有缓冲和无缓冲：
- 无缓冲通道是初始化时没有声明缓冲大小或生命的大小为0
- 无缓冲通道想要传数据时必须要有接收者（可以形象的理解为送快递没有快递柜，必须有人签收）
- 有缓冲通道与无缓冲通道刚好相反（可以形象的理解为送快递有快递柜，可以先放在那，不一定马上需要收）
从通道取值有多种方式：
- 普通的取值: v, ok := <- ch
- for range取值: for i := range ch
- select case: select可以用来监听多个通道，当有通道准备好了则执行对应的case语句
*/
func main() {
	ch := make(chan int, 1) // 有缓冲通道
	ch <- 10                // 将10写入通道
	x := <-ch
	fmt.Println(x)
	close(ch) // 关闭通道

	ch1 := make(chan int) // 无缓冲通道
	// ch1 <- 10
	// fmt.Println("是否可以执行到这") 这样编译可通过，但是运行时报错deadlock，上一行死锁了
	go recv(ch1)
	ch1 <- 10
	fmt.Println("发送成功") // 这样写就没问题，因为先有了接收者，发送者可以发过去了

	ch2 := make(chan int)
	go help(ch2)
	// 第一种取值方式
	for {
		i, ok := <-ch2
		if !ok {
			fmt.Println("结束")
			break
		}
		fmt.Println(i)
	}

	ch3 := make(chan int)
	go help(ch3)
	// 第二种取值方式，ch3关闭后结束
	for i := range ch3 {
		fmt.Println(i)
	}
	fmt.Println("结束")

	// 第三种取值方式
	ch4 := make(chan int)
	go help(ch4)
	for {
		select {
		case i, ok := <-ch4:
			if !ok {
				fmt.Println("通道关闭，正常结束")
				return
			}
			fmt.Println(i)
		default:
			fmt.Println("default")
		}
	}
}

// 无缓冲通道的接收者
func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}

func help(ch chan int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}
