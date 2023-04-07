package main

import "fmt"

type Eatable interface {
	eat()
}

type Apple struct{}

func (a *Apple) eat() {
	fmt.Println("Eating apple. Transforming 5 unit energy from to brain")
}

type Chocolate struct{}

func (a *Chocolate) eat() {
	fmt.Println("Eating Chocolate.")
}

type Robot struct{}

func (r *Robot) GetEnergy(eatable Eatable) {
	//if eatable == "Apple"{
	//	apple := Apple{}
	//	apple.eat()
	//}else if eatable == "Chocolate"{
	//	chocolate := Chocolate{}
	//	chocolate.eat()
	//}
	eatable.eat()
}

func main() {

}
