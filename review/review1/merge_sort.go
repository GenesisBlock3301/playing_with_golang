package main

import (
	"fmt"
	"math"
)

func mergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	mid := int(math.Floor(float64(len(arr) / 2)))
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)

}

func merge(left []int, right []int) []int {
	var mergeList []int
	l1 := len(left)
	l2 := len(right)
	i := 0
	j := 0
	for i < l1 && j < l2 {
		if left[i] < right[j] {
			mergeList = append(mergeList, left[i])
			i++
		} else {
			mergeList = append(mergeList, right[j])
			j++
		}
	}

	if i < l1 {
		mergeList = append(mergeList, left[i:]...)
	}
	if j < l2 {
		mergeList = append(mergeList, right[j:]...)
	}
	return mergeList
}

func main() {
	arr := []int{5, 4, 3, 3, 2, 1}
	fmt.Println(mergeSort(arr))
}
