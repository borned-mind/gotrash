package main
import "fmt"

func half(x float64) (ret float64, divided bool){
	ret = x*0.5
	if int(ret) - int(ret/2)*2 == 0 {
		divided = true
	}
	return
}

func max(x ... float64) (ret float64){
	for _, v := range x{
		if ret < v {
			ret = v
		}
	}
	return
}

func OddGenerator() func () uint{
	i := uint(1);
	return func() (ret uint){
		ret = i;
		i+=2;
		return
	}
}

func fib(x uint) uint{
	if x == 0 {
		 return 0
	}else if x == 1 {
	 	return 1
	}
	return fib(x-1) + fib(x-2)
}

func main(){
	fmt.Println(half(228.4))
	fmt.Println( max(6,4) )
	fmt.Println( max([]float64{1,22,3}...) )
	generator := OddGenerator();
	for i:=0; i<10;i++{
		fmt.Println( generator() );
	}
	var f uint
	for true{
		fmt.Print("Фибоначи: ")
		fmt.Scanf("%d", &f);
		fmt.Println( fib(f) );
	}
}
