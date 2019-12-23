package main

import (
	"encoding/json"
	"fmt"
)

func main()  {

	//对数组的编码
	x:=[5]int{1,3,4,5,7}
	s,err2:=json.Marshal(x)
	if err2!=nil{
		panic(err2)
	}
	fmt.Println(string(s))
}