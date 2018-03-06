package main
import "fmt"


func main(){
//	fmt.Print("Write number operation number: ")
//	var number1,number2 int64
//	var operation string
//	fmt.Scanln(&number1, &operation, &number2)

//	fmt.Println( string(number1) + operation + string(number2) )

	fmt.Println("1+1=", 1+1)
	fmt.Println(len("Hello world"))
	fmt.Println("h"+string("Hello World"[4]) + "lla")
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println(1<<8 | 8 & 0xFF ^ 0xFF)
	fmt.Println(1^2) // xor
	fmt.Println(^228)//not
	fmt.Println(011)
}
