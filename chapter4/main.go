package main

import "fmt"

func main(){
	fmt.Print("Введите имя: ")
	var Имя string 
	//Имя := "Name"
	fmt.Scanln(&Имя)
	Sooqa := 660;
	fmt.Println("Ваше имя " + Имя +"?")
	fmt.Println("Not="+string(^Sooqa))
	fmt.Println("Xor 12="+string(Sooqa^12))
	fmt.Println("And 12="+string(Sooqa&12))
	fmt.Println("Or 12="+string(Sooqa|12))
}
