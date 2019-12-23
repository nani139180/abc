package  main
import (
	"fmt"
)
func main() {
	//多维数组
 //var a [10] int
 //var b [5] int
  //fmt.Println(a,b);	   
  var c [3][4] int
  k:=0;
  for i:= 0; i<3; i++ {
	  for j := 0; j <4; j++	{
		  k++;
		  c[i][j]=k;
		  fmt.Printf("a[%d][%d]=%d,\n",i,j,c[i][j])
		  //frm.fmt.Println(i,j,a[i][j])
	  }
   }

}