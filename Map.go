/*map 映射键到值。
map 在使用之前必须用 make 而不是 new 来创建；
值为 nil 的 map 是空的，并且不能赋值*/
package main
import "fmt"
type Vertex struct {
	Lat, Long float64
}
var m map[string]Vertex
func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
// package main

// import (
// 	"code.google.com/p/go-tour/wc"
// )

// func WordCount(s string) map[string]int {
// 	return map[string]int{"x": 1}
// }

// func main() {
// 	wc.Test(WordCount)
// }
