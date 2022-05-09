package main

import (
	"fmt"
	"strings"
)

func main(){
	var str string = "Nur Amin Sifat"
	split := strings.Split(str," ")
	fmt.Println(len(split))
	for i,v := range split{
		fmt.Println(i,v)
	}
}