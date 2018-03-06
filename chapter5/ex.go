package main
import "fmt"

func EvenGenerator() func(a ... bool) uint{
	i:=uint(0)
	return func(a ... bool) (ret uint){
		if len(a) > 0 && a[0] == true{
			i = 0
		}
		ret = i
		i+=2;
		return
	}
}

func first(){
	fmt.Println("First")
}

func second(){
	fmt.Println("Second")
}

func factorial(x uint) uint{
	if x == 0{
		return 1
	}

	return x * factorial(x-1)
}

func min(s []int) (int){
	min:=s[0]
	for _, value := range s{
		if value < min { min = value }
	}
	return min
}

func main(){
	x := []int{
    		48,96,86,68,
    		57,82,63,70,
    		37,34,83,27,
    		19,97, 9,17,
	}
	fmt.Println(min(x))
	add := func(x ... int) (total int){
		total = 0
		for _, v := range x{
			total += v
		}
		return
	}
	fmt.Println(add(1,2,3,4,5))
	fmt.Println(add(x...))
	Generator := EvenGenerator()
	defer Generator(true)
	for i:=20; i > 0; i--{
		fmt.Print(Generator(), " ")
	}
	fmt.Print("\n")
	fmt.Println(factorial(6))
	defer second()
	first()

}
