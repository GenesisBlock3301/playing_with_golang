package main

import "fmt"

func panicHandler() {
	rec := recover()
	if rec != nil {
		fmt.Println("RECOVER", rec)
	}
}

func employee(name *string, age int) {
	if age > 65 {
		panic("Age cannot be greater than retirement age")
	}

}

func main() {
	empName := "Samia"
	age := 75

	employee(&empName, age)

}
