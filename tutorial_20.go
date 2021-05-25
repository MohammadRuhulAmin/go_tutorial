package main

import (
	"fmt"
	"math"
)

type Geometry interface {
	area() float64
}

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func measure(g Geometry) { //  this is for interface method!
	fmt.Println(g)
}

func main() {

	r := Rectangle{
		width:  2.5,
		height: 3.5,
	}
	fmt.Println("The Area of Rectangle is : ", r.area())

	c := Circle{
		radius: 5.63,
	}
	fmt.Println("The Area of the Circle is : ", c.area())
	measure(r)
	measure(c)

}
