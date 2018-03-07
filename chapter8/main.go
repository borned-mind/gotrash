package main
import "fmt"

func getAdressAndValue(xPtr *int){
	fmt.Println( "Adress: ", xPtr, "Value: ", *xPtr )
}

func zero(xPtr *int){
	getAdressAndValue(xPtr)
	set(xPtr,0)
}

func square(x * float64){
	*(x) = *(x) * *(x)
}

func set(xPtr *int, set int){
	*xPtr=set
}

func swap(a *int, b *int){
	tmp := *(a);
	*(a) = *(b);
	*(b)=tmp;
}

func main(){
	x := 5;
	zero(&x)
	xPtr := new(int)
	set(xPtr, 1)
	getAdressAndValue(xPtr)
	x=1;
	y:=2;
	swap(&x,&y);
	fmt.Println(x,y);
	c := new(float64)
	*(c)=1.5
	square(c)
	fmt.Println(*c)
//	fmt.Println(x)
}
