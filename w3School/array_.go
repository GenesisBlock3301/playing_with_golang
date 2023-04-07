package main

import "fmt"

func testCount(val int) int {
	if val == 11 {
		return 0
	}
	println(val)
	return testCount(val + 1)
}
func main() {

	type Person struct {
		name string
		age  int
		//job    string
		//salary int
	}
	var population []Person
	var name string
	var age int
	for i := 0; i < 3; i++ {
		fmt.Println("Enter name:")
		_, err := fmt.Scanln(&name)
		if err != nil {
			return
		}
		fmt.Println("Enter age:")
		_, err = fmt.Scanln(&age)
		if err != nil {
			return
		}
		p := Person{name: name, age: age}
		population = append(population, p)
	}
	fmt.Println(population)

}
