package main

import "fmt"

func countChar(str string) []int {
	count := make([]int, 128)
	for _, val := range str {
		count[val]++
	}
	return count
}
func longestPalindrom(str string) int {
	alpha := countChar(str)
	n := len(str)
	odd_count := 0
	for _, val := range alpha {
		if val%2 != 0 {
			odd_count++
		}
	}
	if odd_count > 1 {
		return n - (odd_count - 1)
	} else {
		return n
	}
}
func main() {
	s := "abccccdd"
	//fmt.Println(countChar(s))
	fmt.Println(longestPalindrom(s))
}
