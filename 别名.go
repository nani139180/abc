package main

import "fmt"

type bigint int64
type (
	long int64
	char byte
)
var ch char='b'
var b long=111
var a int64
func main() {

	fmt.Printf("a type is %T\n",a)

}
