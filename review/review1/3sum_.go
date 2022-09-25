package main

import (
	"reflect"
	"sort"
)

func checkExistence(checkedArray []int, arr [][]int) bool {
	for _, v := range arr {
		if reflect.DeepEqual(v, checkedArray) {
			return true
		}
	}
	return false
}

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	var arr [][]int
	for i := 0; i < n-2; i++ {
		//skip duplicates
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, h := i+1, n-1
		for l < h {
			switch sum := nums[i] + nums[l] + nums[h]; {
			case sum == 0:
				arr = append(arr, []int{nums[i], nums[l], nums[h]})
				l++
			case sum < 0:
				l++
			case sum > 0:
				h--
			}
		}
	}
	return arr
}

func main() {
	//nums := []int{-1, 0, 1, 2, -1, -4}
	//arr := [][]int{
	//	{5, 6, 7, 8},
	//	{1, 2, 3},
	//}
	//sort.Sort(arr[0] > arr[1])

	//fmt.Println(arr[0] > arr[1])
}
