package main

import "fmt"
//channel 是有类型的管道，可以用 channel 操作符 <- 对其发送或者接收值。
//go关键字可以用来开启一个goroutine(协程))进行任务处理，而多个任务之间如果需要通信
func main() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
}