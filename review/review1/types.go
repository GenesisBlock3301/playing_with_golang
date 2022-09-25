package main

import "fmt"

func numbers() {
	fmt.Println("1+1 = ", 1+1)
	fmt.Println("1.0+1.0 = ", 1.0+1.0)
	fmt.Println(len("Nur amin sifat"))
	fmt.Println("Nur Amin Sifat"[1])
	fmt.Println(true && true)
}

func main() {
	numbers()
}
