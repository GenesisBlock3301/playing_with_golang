package main

import "fmt"

type IShape interface {
	DrawSquare() float64
	DrawRectangle() float64
	DrawCircle() float64
}

type Circle struct{}

func (receiver *Circle) DrawSquare() {
	fmt.Println(64)
}

func main() {
	c := Circle{}
	c.DrawSquare()
}
