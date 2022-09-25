package main

import "fmt"

func makeGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 1
		return
	}
}

func main() {
	nextVal := makeGenerator()
	fmt.Println(nextVal())
	fmt.Println(nextVal())
}
