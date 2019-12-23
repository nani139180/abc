package main

import "fmt"

//import "fmt";
type  pss struct {
	pers string;
   age int;
}
type haha struct {
	//pss;
	name string
	age int
}

func main()  {
	//haha :=haha{pss{"ddd",17},"ddd",33}
	haha:=haha{"dd",19}
    fmt.Println(haha.name,haha.age);
}
