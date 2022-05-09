package main

import "fmt"

type person struct{
	name string
	age float32
}

func newPerson(name string,age float32) *person{
	p := person{name:name,age: age}
	return &p
}
func main(){


	fmt.Println(newPerson("N A Sifat",26))
}