// Program to find area and perimeter using interface


package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}


type rectangle struct {
	width, height float64
}
type circle struct{
	radius float64
}


func (c circle) area()float64{
	return math.Pi * math.Pow(c.radius,2)
}

func (c circle) perimeter() float64{
	return 2 * math.Pi * c.radius
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.height + r.width)
}

func printShape(s shape){
	fmt.Printf("Shape:%v\n",s)
	fmt.Printf("area:%v\n",s.area())
	fmt.Printf("perimeter:%v\n",s.perimeter())
}

func main(){
	c1 := circle{5.}
	r1 := rectangle{width: 3.,height: 2.1}

	printShape(c1)
	printShape(r1)
}
