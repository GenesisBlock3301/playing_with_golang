package main

import (
	"fmt"
	"regexp"
)

func main(){
	re := regexp.MustCompile(".com")
	// exist or not
	fmt.Println(re.FindString("googl.com"))

	// Find index
	fmt.Println(re.FindStringIndex("google.com"))

	n_re := regexp.MustCompile("f([a-z]+)ing")
	fmt.Println(n_re.FindStringSubmatch("flying"))
}