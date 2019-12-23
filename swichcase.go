package main

import (
	"fmt"
	)

func main() {
	/*
	var score=2
	switch score {
	case 1:
		fmt.Println("这个数字是1")
	case 2:
		fmt.Println("这个数字是2")
	case 3:
		fmt.Println("这个数字是3")
	}
	*/
	//var score=2
	switch num:=2;num{
	case 1:
		fmt.Println("这个数字是1")
	case 2:
		fmt.Println("这个数字是2")
	case 3:
		fmt.Println("这个数字是3")
	default:
	}
}
