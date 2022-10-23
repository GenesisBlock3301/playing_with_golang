package main

import "fmt"

func findMaxVal(arr []int) int {
	max := arr[0]
	for _, i := range arr {
		if i > max {
			max = i
		}
	}
	return max
}
func countingSort(arr []int) []int {
	maxVal := findMaxVal(arr) + 1
	size := len(arr)
	count := make([]int, maxVal)
	var output []int
	// count element
	for i := 0; i < size; i++ {
		count[arr[i]]++
	}
	for i, val := range count {
		if val > 0 {
			addingArray := makeArray(i, val)
			output = append(output, addingArray...)
		}
	}

	return output
}
func makeArray(index int, value int) []int {
	var a []int
	i := 0
	for i < value {
		a = append(a, index)
		i++
	}
	return a
}
func main() {
	arr := []int{4, 2, 2, 8, 3, 3, 1, 11}

	fmt.Println(countingSort(arr))
}
