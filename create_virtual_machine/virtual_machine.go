package main

import (
	"fmt"
	"strconv"
)

func virtualMachine(program []string) {
	// `programCounter` keep the track of current instruction
	// `stackPointer` manage the stack
	var programCounter, stackPointer int
	stack := make([]int, len(program))

	for programCounter < len(program) {
		currentInstruction := program[programCounter]

		switch currentInstruction {
		case "PUSH":
			val := program[programCounter+1]
			value, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			stack[stackPointer] = value
			stackPointer++
			programCounter++
		case "ADD":
			right := stack[stackPointer-1]
			stackPointer--
			left := stack[stackPointer-1]
			stackPointer--
			stack[stackPointer] = left + right
			stackPointer++
		case "MINUS":
			right := stack[stackPointer-1]
			stackPointer--
			left := stack[stackPointer-1]
			stackPointer--
			stack[stackPointer] = left - right
			stackPointer++
		}
		programCounter++
	}
	fmt.Println("stack-top: ", stack[stackPointer-1])
}

func main() {
	program := []string{
		"PUSH", "3",
		"PUSH", "4",
		"ADD",
		"PUSH", "5",
		"MINUS",
	}
	virtualMachine(program)
}
