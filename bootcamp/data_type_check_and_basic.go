package main

import "fmt"

func main() {
	var sname, age, isStudent = "Sifat", 26, false
	fmt.Println(sname, age, isStudent)

	fmt.Printf("%v = %T\n", sname, sname)
	fmt.Printf("%v = %T\n", age, age)
	fmt.Printf("%v = %T", isStudent, isStudent)
}
