package main

import (
	"fmt"
	"sort"
)

func main() {
	var arr []int
	var n int
	fmt.Scanln(&n)
	for i := 1; i < n+1; i++ {
		var val int
		fmt.Scanln(&val)
		arr = append(arr, val)
	}
	//found := false
	sort.Ints(arr)
	fmt.Println(arr)
}
