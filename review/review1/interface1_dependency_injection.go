package main

import "fmt"

type animal interface {
	breathe()
	walk()
}

type lion struct {
	age int
}

func (l lion) breathe() {
	fmt.Println("Lion can breathe")
}

func (l lion) walk() {
	fmt.Println("Lion can walk..")
}

type dog struct {
	age int
}

func (l dog) breathe() {
	fmt.Println("dog can breathe")
}

func (l dog) walk() {
	fmt.Println("dog can walk..")
}

func main() {
	var a, d animal
	a = lion{age: 40}
	callBreath(a)
	callWalk(a)
	d = dog{age: 10}
	callBreath(d)
	callWalk(d)
}

func callBreath(a animal) {
	a.breathe()
}

func callWalk(a animal) {
	a.walk()
}
