package main

import (
	"fmt"
	"math"
)

// why use and need structs and interfaces

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1

	return math.Sqrt(a*a + b*b)
}

type Circle struct {
	x, y, r float64
}

func circleAEW(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

//special type of function call method/ receiver function
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func main() {
	//var c Circle
	//c := new(Circle)
	c := Circle{1, 2, 3}
	fmt.Println(c.area())

}
