package main

import "fmt"

func start(val int) {
	for i := 0; i < val; i++ {
		fmt.Println(val)
	}
}
func main() {
	fmt.Println("Start")
	a := 5
	go start(a)
	b := 100
	fmt.Println("End")
	fmt.Println(b)
}
