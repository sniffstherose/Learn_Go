package main

/*
// 导入包的写法
1.
import "fmt"
import "os"
2. 常用
import (
	"fmt"
	"os"
)
*/
/*
.操作
import . "fmt"	// 调用Println可以不写fmt了，但是不建议，防止出现同名情况
func main() {
	Println("test")
}
*/

/*
给包起别名
import io "fmt"
func main() {
	io.Println("test")
}
*/

/*
忽略包，_操作
import _ "fmt"	// 不使用也不会报错，但是也不能使用。。。
func main() {
	fmt.Println("test")
}
*/