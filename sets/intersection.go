package main

import "fmt"

func Intersection(a []int, b []int) []int {
	temp := make(map[int]bool)
	for _, val := range b {
		temp[val] = true
	}
	var insect []int
	for _, val := range a {
		if _, ok := temp[val]; ok {
			insect = append(insect, val)
		}
	}
	return insect
}

func main() {
	var a = []int{1, 2, 3, 4, 5}
	var b = []int{2, 3, 5, 7, 11}
	fmt.Println(Intersection(a, b))
}
