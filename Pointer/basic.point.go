package main

import (
	"fmt"
)
// pointer variable hold address of another variable
func zero(x *int){
	*x = 0
}

func main(){
	x := 9
	zero(&x)
	fmt.Println((x))
}
