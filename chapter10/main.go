package main
import (
	"fmt"
	"time"
	"math/rand"
	)
var c uint
func f(n int){
	fmt.Println("cycle",c)
	c+=1
	for n=n; n>0; n--{ // n=n for error, not used -_-
		fmt.Println(n)
		amt := time.Duration( rand.Intn(250) )
		time.Sleep( time.Millisecond * amt )
	}

}


func main(){
	for i:=10;i>0;i--{
		go f(i)
	}
	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
}
