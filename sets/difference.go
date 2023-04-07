package main

import "fmt"

func Difference(a []int, b []int) []int {
	isCommon := make(map[int]bool)
	for _, val := range b {
		isCommon[val] = true
	}
	var diff []int
	for _, val := range a {
		if _, ok := isCommon[val]; !ok {
			diff = append(diff, val)
		}
	}
	return diff
}
func main() {
	var a = []int{1, 2, 3, 4, 5}
	var b = []int{2, 3, 5, 7, 11}
	fmt.Println(Difference(a, b))
}
