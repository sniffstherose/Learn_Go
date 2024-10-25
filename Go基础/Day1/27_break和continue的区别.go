package main

import "fmt"
import "time"

func main() {
	i := 0;
	for {	// for后面不跟条件，则条件一定为真相当于c中的while(1)
		i++
		time.Sleep(time.Second)
		if (i == 5) {
			// break	// 跳出整个循环
			continue 	// 跳出单次循环
		}
		fmt.Println(i)
	}
}