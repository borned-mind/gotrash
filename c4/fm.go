package main

import "fmt"

func funtToMetr(faring float64) (float64){
   return faring*0.3048

}

func main(){
   var far float64
   fmt.Print("Введите футы: ");
   fmt.Scanf("%f", &far)
   fmt.Println("В метрах: ", funtToMetr(far))
}

