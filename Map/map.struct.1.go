package main

import (
	"fmt"
)

type contract struct {
	email string
	phone string
}
type student struct {
	firstName string
	lastName  string
	contract
}

// This is receiver function like OOP based class method
func (s student) print() {
	fmt.Println("First name:", s.firstName, "Last Name:", s.lastName)
}

func main() {
	// sifat := student{firstName: "N A ", lastName: "Sifat",
	// 	contract: contract{email: "nas@gmail.com", phone: "+880809094353"}}
	

}
