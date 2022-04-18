package main

import (
	"fmt"
)

func main(){
	var a int = 45
	// pointer variable hold the address of variable
	var s *int = &a
	fmt.Println((s))
}