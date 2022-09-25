package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, his name is", p.Name)
}

type Android struct {
	Person Person
	Model  string
}

func (a *Android) detail() map[string]string {
	return map[string]string{"name": a.Person.Name, "model": a.Model}
}

func main() {
	a := new(Android)
	a.Person.Name = "Android user"
	a.Model = "A5"
	fmt.Println(a.detail())
}
