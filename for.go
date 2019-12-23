package main
import "fmt"
func main() {
	str:="abd"
	for i:=0;i<len(str);i++{
		fmt.Printf("str[%d]=%c\n",i,str[i])
	}
}
