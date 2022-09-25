package main

import "fmt"

func findIndex(searchVal int, nums2 []int) int {
	for i, val := range nums2 {
		if val == searchVal {
			return i
		}
	}
	return -1
}

func nextGreater(nums1 []int, nums2 []int) []int {
	var result []int
	var max int
	for _, val := range nums1 {
		index := findIndex(val, nums2)
		if index+1 <= len(nums2)-1 {
			if nums2[index+1] > nums2[index] {
				result = append(result, max)
			} else {
				max = -1
			}
		} else {
			result = append(result, -1)
		}
	}
	return result
}

func main() {
	nums1 := []int{4, 1, 2}
	nums2 := []int{3, 1, 4, 2}

	fmt.Println(nextGreater(nums1, nums2))
}
