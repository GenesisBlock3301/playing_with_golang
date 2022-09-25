package main

import "fmt"

func loop1() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i += 1
	}
}

func loop2() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func main() {
	fmt.Println("While like loop")
	loop1()
	fmt.Println("For like loop")
	loop2()
}
