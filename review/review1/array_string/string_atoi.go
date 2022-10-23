package main

import (
	"fmt"
	"math"
	"strings"
)

func myAtoi(s string) int {
	result := 0
	n := len(s)
	if n == 0 {
		return 0
	}
	s = strings.TrimSpace(s)
	mul := 1
	start := 0
	if string(s[0]) == "-" {
		mul = -1
		start = 1
	} else if string(s[0]) == "+" {
		start = 1
	}
	for i := start; i < n; i++ {
		if !(s[i] >= 48 && s[i] <= 57) {
			return result * mul
		}
		result = (result * 10) + (int(s[i]) - 48)
		//var minInt32 int32
		minInt32 := math.Pow(2, 31)
		if result*mul <= -int(minInt32) {
			return int(-minInt32)
		} else if result*mul >= int(minInt32-1) {
			return int(minInt32 - 1)
		}
	}
	return result * mul
}

func main() {
	s := "-91283472332"
	fmt.Println(myAtoi(s))
}
