package main

import "fmt"



func Double3(a int, b int, c int) (int,int,int){
	return a*2, b*2, c*2
}

func init3() (int8){
	var(
		a=3
		b=4
		c=5
	)
	const (
		a1=6
		b1=7
		c2=8
	)
	return int8(a+b+c+a1+b1+c2)
}
/*
func testLink(test *int){
	*test = 228+228
}
*/


func main(){
	test1, test2, test3 := Double3(6,3,2)

	fmt.Println(test1,test2,test3)
	fmt.Print("Write 3 numbers: ")
	fmt.Scanf("%d %d %d",&test1,&test2,&test3) // pointers

	test1,test2,test3 = Double3(test1,test2,test3)
	fmt.Println(test1,test2,test3)
//	testLink(&test3)
	main()
}
