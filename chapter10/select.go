//piPES?

package main

import ("fmt";"time")


func main(){
/*
i := 5
switch i{
	case 1:
	case 2:
	case 3:
	case 4:
	case 5:
		fmt.Println("not okey:(")
	case 6:
		fmt.Println("okey")
}
*/

//time.Sleep(time.Second*5)

c1:=make(chan string)
c2:=make(chan string)

go func(){
	for{
		c1<-"ping"
		time.Sleep(time.Second*2)
	}
}();

go func(){
	for{
		c2<-"pong"
		time.Sleep(time.Second*1)
	}
}();


go func(){
	for{
		select{
			case msg1:=<-c1:
				fmt.Println(msg1)
				break;
			case msg2:=<-c2:
				fmt.Println(msg2)
				break;
			case <- time.After(time.Second):
				fmt.Println("Time out")
//			default:
//				fmt.Println("not anything")
		}
	}
}();


var input string;
fmt.Scanln(&input)


}
