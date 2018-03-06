package main

import "fmt"


func main(){
	for i:=1; i<=10;i++ {
		if i % 2 == 0{
			fmt.Println("even")
		}else{
			fmt.Println("odd")
		}
		switch(i){
			case 0: fmt.Println("Zero");break;
			case 1: fmt.Println("One");break;
			case 2:fmt.Println("Two");break;
			case 3:fmt.Println("Three");break;
			case 4:fmt.Println("Four");break;
			case 5:fmt.Println("Five");break;
			default:fmt.Println("Unknown");break;
		}
	}
	i:=100
	for i>0{
		if( i % 3 == 0){ fmt.Print("Fizz") }
		if( i % 5 == 0){ fmt.Print("Buzz") }
		fmt.Println("\n");
		i--
	}
	//for !false{} // 
	//for true{} // 

	//for i:=10;i--;{} // wrong
}
