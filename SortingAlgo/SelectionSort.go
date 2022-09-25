package main

import (
	"fmt"
)

func selection(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		indexMin := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[indexMin] {
				indexMin = j
			}
		}
		if indexMin != i {
			arr[i], arr[indexMin] = arr[indexMin], arr[i]

		}
	}
	return arr

}

func main() {
	var size int
	fmt.Printf("Enter your array size: ")
	fmt.Scanln(&size)
	var arr = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	fmt.Println("Sorted array: ", selection(arr))

}
