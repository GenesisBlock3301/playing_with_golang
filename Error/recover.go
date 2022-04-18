package main

import (
	"fmt"
)

func SaveDivide(num1, num2 int) int{
	defer func ()  {
		fmt.Println(recover())
	}()
	quotient := num1/num2
	return quotient
}

func main(){
	fmt.Println(SaveDivide(10,0))
	fmt.Println(SaveDivide(10,1))
}