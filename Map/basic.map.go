package main

import (
	"fmt"
)

func main(){
	x := map[int]string{1:"Sifat",2:"Nas"}
	fmt.Println(x[1])
	// Map insert & update
	m := make(map[string]int)
	fmt.Println(m)
	m["key1"] = 10
	m["key2"] = 2
	m["key3"] = 50
	fmt.Println(m)
	// retrieve element
	ele,ok := m["key3"]
	fmt.Println(ele,ok)
	// map delete
	delete(m,"key3")
	fmt.Println(m)

}