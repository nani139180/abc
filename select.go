package main
import (
	"fmt"
	//"time"
)
//select
//select 语句使得一个 goroutine 在多个通讯操作上等待。
//select 会阻塞，直到条件分支中的某个可以继续执行，这时就会执行那个条件分支。当多个都准备好的时候，会随机选择一个。
func main() {
	ch:=make(chan int)
	select {
	case <-ch:
		fmt.Print("DDDD")
	default:
		fmt.Print("saaa")
	}
}