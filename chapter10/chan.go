package main

import ("fmt";"time";"math/rand")

func pinger(c chan <- string){
	for true{
		c <- "ping"
	}
}

func printer(c <- chan string){
	for{
		msg := <-c
		fmt.Println( msg, time.Now() )
		time.Sleep(time.Second)
	}
}

func test_input(t chan <- int){
	for{
		t <- rand.Intn(2000)
		time.Sleep(5*time.Second)
	}
}

func test_output(t <- chan int){
	for{
		fmt.Println(<-t)
	}
}

func main(){

 c := make(chan string)
 t := make(chan int)

 go test_input(t)
 go test_output(t)
//
 go pinger(c)
 go printer(c)
 var input string
 fmt.Scanln(&input)

}
