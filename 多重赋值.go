package main
import "fmt"
a, b := 10, 20
a, b = b, a
func test()(int int){
	return 1,2
}
func main() {
	a,b:=test()
}