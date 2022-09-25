package main

import "fmt"

var x string = "Hello-world"

func const_example() string {
	const x string = "Const test"
	y := "changex"
	return y
}
func main() {
	test()
	fmt.Println(const_example())
}

func test() {
	fmt.Println(x)
}
