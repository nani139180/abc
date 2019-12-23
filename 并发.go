package main

import (
	"fmt"
		"strconv"
)
func Read(ch chan  int){
	value:=<-ch
	//fmt.PrintIn(strconv.itoa(value))
	fmt.Println(strconv.Itoa(value))
}
func Write(ch chan int){
	ch<-10
}
func main() {
	ch:=make(chan int)
	go Read(ch)
	go Write(ch)
	//time.Sleep(10)
	for _,v:=range chs{
		<-v
	}
}