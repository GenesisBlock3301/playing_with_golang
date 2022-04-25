package main

import (
	"errors"
	"fmt"
)

// Interface for Error handle
type error interface {
	Error() string
}

// Struct for zero division
type DivZero struct{}

//
func (myerr *DivZero) Error() string {
	return "Cannot divide by zero!"
}
func divide(x, y int) (int, error) {
	if y == 0 {
		return -1, errors.New("Cannot divid by zero!")
	}
	return x / y, nil
}

func main() {
	answer, err := divide(5, 5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(answer)
	}
	myerr := &DivZero{}
	fmt.Println(myerr)
}
