package main
import "fmt"

func EvenGenerator() func(a ... bool) uint{
	i:=uint(0)
	return func(a... bool) (ret uint){
		ret = i
		i+=2
		if len(a) > 0 && a[0] == true{
			i = 0
		}
		return
	}
}


func main(){
	generator := EvenGenerator()
	for i:=10;i>0;i--{
		fmt.Println(generator())
	}

}
