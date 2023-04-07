package main

import "fmt"

// Animal The SRP states that
//type Animal struct {
//	Name string
//}
//
//func (a *Animal) GetName() string {
//	return a.Name
//}
//
//func animalSound(animals []Animal) {
//	for _, animal := range animals {
//		if animal.Name == "lion" {
//			fmt.Println("Roar")
//		} else if animal.Name == "mouse" {
//			fmt.Println("squeak")
//		}
//	}
//}

type Animal interface {
	getName() string
	makeSound() string
}

type Lion struct {
	Name string
}

func (l *Lion) getName() string {
	return l.Name
}
func (l *Lion) makeSound() string {
	return "Roar"
}

type Mouse struct {
	Name string
}

func (m *Mouse) getName() string {
	return m.Name
}
func (m *Mouse) makeSound() string {
	return "Squawk"
}

func getAnimalSound(animals *[]Animal) {
	for _, animal := range *animals {
		fmt.Println(animal.makeSound())
	}
}

func main() {
	var animals []Animal
	animals = append(animals, &Lion{Name: "lion"}, &Mouse{Name: "mouse"})
	getAnimalSound(&animals)
}
