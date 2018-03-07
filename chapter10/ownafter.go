package main
import ("fmt";"time")

func sleep(t uint){
	fmt.Println(time.Now())
	s := <- time.After( time.Duration(t)*time.Second);
	fmt.Println(s)
}

func main(){
	sleep(5)
}
