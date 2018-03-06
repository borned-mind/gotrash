package main
import "fmt"

func main(){

//	x:= make([]int,5, 10)
	slice:=[]int{1,2,3}
	x2 := append(slice,227,228)
	x3 := slice[:3]
	x4:=slice[:]
	x5:=slice[0:]
	fmt.Println(x2,x3,x4,x5)


}
