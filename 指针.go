package main
import ("fmt")
func swap(p1 * int,p2 * int){
	*p1,*p2 =*p2,*p1;
}
func main() {
  a :=10;
 var p * int;p=&a;
 fmt.Printf("p=%d,&a=%v,*p=%v\n",p,&a,*p)
  a,b :=10,20
  swap(&a,&b)
  fmt.Printf("a=%d,b=%d\n",a,b)
}