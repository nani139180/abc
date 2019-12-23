package main
import (
	"fmt"
)

func main() {
	var slice []int	 // 定义切片
	fmt.Printf("%d \n", slice)  // 打印切片地址  切片的变量名本身就是一个地址,不需要用&取地址   0x0  空切片的地址
	slice = append(slice, 1,2,3) // append为切片添加数据时，切片的地址可能会改变(切片后面如果有足够的连续空间时，地址就不会改变)
	fmt.Printf("%d \n", slice)  // 0xc00004

	sq1 :=[] int{1,2,3}
	fmt.Println("sql=",sq1)
}