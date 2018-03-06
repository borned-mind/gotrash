package main

import (
	"fmt"
       )


func main(){
	var name string
	fmt.Print("My name is?: ")
	fmt.Scanln(&name)
	fmt.Println("Hello my name is " + name)
}
