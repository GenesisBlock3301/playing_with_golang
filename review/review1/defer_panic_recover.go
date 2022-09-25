package main

import "fmt"

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}

func panicRecover() {
	defer func() {
		str := recover()
		fmt.Println(str, "999")
	}()
	panic("PANIC")
}

func main() {
	defer second()
	first()
	panicRecover()
}
