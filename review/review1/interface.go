package main

import (
	"fmt"
	"math"
)

type tank interface {
	Tarea() float64
	Volume() float64
}

type myValue struct {
	radius float64
	height float64
}

func (m *myValue) Tarea() float64 {
	return m.radius * m.height
}

func (m *myValue) Volume() float64 {
	return math.Pi * m.radius * m.radius * m.height
}

func main() {
	var t tank
	t = &myValue{10, 14}
	fmt.Println(t.Tarea(), t.Volume())
}
