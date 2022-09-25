package main

import "fmt"

type iBase interface {
	say()
}

type base struct {
	color string
}

func (b *base) say() {
	b.clear()
}

func (b *base) clear() {
	fmt.Println("Clear from base function")
}

type child struct {
	base
	style string
}

func check(b iBase) {
	b.say()
}

func (b *child) clear() {
	fmt.Println("Clear from child class")
}
func main() {
	base := base{color: "Red"}
	child := &child{
		base:  base,
		style: "somestyle",
	}
	child.say()
	//fmt.Println(child.base.color)
	//check(child)
}
