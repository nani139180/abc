package main

import (
	"fmt"
	"time"
)

func test_Rou(){
	fmt.Println("ddd")
}
func Add (x,y int){
	z:=x+y
	fmt.Println(z)
}
func main()  {	
	go test_Rou()
	for index := 0; index < 10; index++ {
		go Add(index,index)
	}
	time.Sleep(10)
}
