package main

import (
	"fmt"
)

func main(){
	var m map[string]int = map[string]int{
		"apple":5,"pear":7,"orange":9,
	}
	val, ok := m["apple"]
	fmt.Println(val,ok)
	// mp := make(map[string]int)
	// fmt.Println(mp)

}