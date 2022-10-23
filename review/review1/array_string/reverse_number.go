//https://leetcode.com/problems/palindrome-number/description/
package main

import "fmt"

func reverseNumber(x int) int {
	result := 0
	for x > 0 {
		rem := x % 10
		result = result*10 + rem
		x = int(x / 10)
	}
	return result
}

func isPalindrome(x int) bool {
	if reverseNumber(x) == x {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(reverseNumber(122))
}
