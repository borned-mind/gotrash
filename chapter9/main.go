package main

import (
	"fmt";"math"
       )

type Person struct{
	name string
}

func (p *Person) talk() *Person{
	fmt.Println("Hello, i am " + p.name)
	return p
}

func (p *Person) setName(name string) *Person{
	p.name = name
	return p
}

type Robo struct{
	Person
	Model string
}

type Circle struct{
	x, y, r float64
}

type Square struct{
	x1,x2 float64
	y1,y2 float64
}

func (c *Circle) perimetr() float64{
	// pi = area/diametr
	// diametr = 2*r
	// area = pi*diametr

	//d = pi/area = d*pi/pi = d

	// c = a/b
 	// a = b*c
	// b = a/c

	// c = a*b
	// a = c/b
	// b = c/a

	return math.Pi * c.r * 2
}


func (r *Square) area() float64{ // square here as rectangle
	a:=r.x2-r.x1;
	b:=r.y2-r.y1;
	return a*b
}

func (r *Square) perimetr() float64{
        a:=r.x2-r.x1;
        b:=r.y2-r.y1;
	return a*2+b*2
}

func (c *Circle) area() float64{
	//if Circle expand to triangle then we will see triangle with H=radius and base=area
	//then area of this triangle is 0.5 * base*H
	/*
		1/2*perimetr*radius
		1/2 * pi * r * 2 * r = 1/2*2 * r * r * pi = 2/2 * r^2 * pi = r^2pi
	*/
	defer func(){
		str := recover()
		if( str != nil ){
			fmt.Println(str)
		}
	}()
	if !(c.r*c.r*math.Pi == 0.5*c.perimetr()*c.r){
		fmt.Println(c.r*c.r*math.Pi,  0.5*c.perimetr()*c.r)
		panic("Error xD")
	}
	return c.perimetr()*c.r
}

type Shape interface{
	area() float64
	perimetr() float64
}

func fArea(shapes ... Shape) (area float64){
	for _, v := range shapes{
		area += v.area()
	}
	return
}

func fPerimetr(shapes ... Shape) (per float64){

	for _, s := range shapes{
		per += s.perimetr()
	}
	return
}


func main(){
	var P Person
	P.setName("Human").talk()
	t := new(Robo)
	t.Person.setName("Android").talk()

	c := Circle{1,2,3}
	s := new(Square)

	s.x1=3;s.x2=6;s.y1=5;s.y2=5;

	fmt.Println(fArea(s,&c))
	fmt.Println( fPerimetr(s, &c) )
}
