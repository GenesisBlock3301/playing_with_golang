package main

import (
	"fmt"
	"regexp"
)

func main() {
	inputText := "I love Bangladesh"
	match, error := regexp.MatchString("[A-z]desh", inputText)
	if error == nil {
		fmt.Println("Match", match)
	} else {
		fmt.Println("Error", error)
	}
}
