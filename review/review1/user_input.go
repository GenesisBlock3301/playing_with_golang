package main

import "fmt"

func example(x int64) int64 {
	return x + 1
}
func main() {
	fmt.Println("Enter your number: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println(output)
	fmt.Println("Enter second number: ")
	var x int64
	fmt.Scanf("%d", &x)
	fmt.Println(example(x))
}
