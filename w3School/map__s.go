package main

import "fmt"

func main() {

	a := map[int]map[string][]int{
		0: {
			"id": {30, 3, 4, 5},
		},
		1: {
			"id": {30, 3, 4, 5},
		},
	}
	fmt.Println(a[0]["id"][0])

}
