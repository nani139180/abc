package main;

import "fmt";

func main()  {
	// x :=1
	//for x<=100 {
	//	x++
	//}
	//fmt.Println(x)
	x:=2;
	switch x {
	case 1:
		fmt.Println("1");
	case 2:
		fallthrough;
	case 3:
		fmt.Println("3")
	}
}
