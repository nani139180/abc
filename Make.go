package main

import "fmt"

func main()  {
	//var  a  =make(map[string]int);
   //test(1,2,"sada")
   fmt.Println(test(1,2,"sada"));
}
func  test(a ,b int ,c string) (int,int,string) {
	return a,b,c;
}
