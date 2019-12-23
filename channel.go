package main

import (
	"time"
	"fmt"
)

func test_channel()  {
	//ch:=make(chan int)
	ch<-1
	ch<-1
	fmt.Println("come to end go runtime 1")
}

var  ch  chan  int
func main()  {
	ch =make(chan int,0)
	go test_channel()
	time.Sleep(2* time.Second)
	fmt.Println("runing end")
	<-ch
	time.Sleep(time.Second)
}