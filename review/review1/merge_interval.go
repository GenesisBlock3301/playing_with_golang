package main

import (
	"fmt"
	"sort"
)

func mx(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func mergeInterval(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] < intervals[j][1] {
			return true
		}
		if intervals[i][1] == intervals[j][1] && intervals[i][0] < intervals[j][0] {
			return true
		}
		return false
	})
	n := len(intervals)
	output := [][]int{intervals[0]}
	for i := 1; i < n; i++ {
		l := len(output)
		if output[l-1][1] >= intervals[i][0] {
			output[l-1][1] = mx(output[l-1][1], intervals[i][1])
		} else {
			output = append(output, intervals[i])
		}
	}
	return output
}

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(mergeInterval(intervals))

}
