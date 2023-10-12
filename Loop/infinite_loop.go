package main

import "fmt"

func main() {
	var a, b float32
	var X int
	for {
		for {
			fmt.Scanln(&a)
			if a <= 10 && a >= 0 {
				break
			} else {
				fmt.Println("nota invalida")
			}
		}
		for {
			fmt.Scanln(&b)
			if b <= 10 && b >= 0 {
				break
			} else {
				fmt.Println("nota invalida")
			}
		}
		fmt.Println("media =", (a+b)/2)
		for {
			fmt.Scanln(&X)
			if X == 1 || X == 2 {
				fmt.Println("novo calculo (1-sim 2-nao)")
				break
			} else {
				fmt.Println("novo calculo (1-sim 2-nao)")
			}
		}
		if X == 2 {
			break
		}
	}
}
