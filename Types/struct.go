package main

import (
	"fmt"
	"reflect"
)

type employee struct{
	person
	emId int
}

type person struct{
	firstname string
	lastname string
	age int
}

func (p person) details(){
	fmt.Println("I am a person.",reflect.TypeOf(p))
}

func (e employee) details(){
	fmt.Println("I am a employee and my id",e.emId,"name is",e.person.firstname,e.person.lastname)
}
func main(){
	p1 := person{age:30,firstname: "N A",lastname: "Sifat"}
	p1.details()
	e1 := employee{person: p1,emId: 1}
	e1.details()

}