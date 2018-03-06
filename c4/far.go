package main

import "fmt"

func farToCel(faring float64) (float64){
	return (faring - 32) * 5/9

}

func main(){
	var far float64
	fmt.Print("Введите температуру в фарингейтах: ");
	fmt.Scanf("%f", &far)
	fmt.Println("В Цельсиях: ", farToCel(far))
}
