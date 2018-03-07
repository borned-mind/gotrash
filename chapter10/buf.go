package main
import "fmt"



func main(){
	//c:=make(chan int,20)
	t:=make(chan int, 1)
	for i:=5;i>0;i--{
		t<-5
		fmt.Println("putted")
		fmt.Println(<-t)
	}
	//now error
	for{
		t<-5
	}

}

