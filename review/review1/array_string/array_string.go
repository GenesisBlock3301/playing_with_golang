package main

import (
	"fmt"
	"reflect"
)

func isExist(str string, counter map[string]int) bool {
	for key, _ := range counter {
		if str == key {
			return true
		}
	}
	return false
}

func Counter(str string) map[string]int {
	counter := make(map[string]int)
	for _, val := range str {
		if isExist(string(val), counter) {
			counter[string(val)] += 1
		} else {
			counter[string(val)] = 1
		}
	}
	return counter
}

func removeZero(window map[string]int) map[string]int {
	for key, _ := range window {
		if window[key] == 0 {
			delete(window, key)
		}
	}
	return window
}

func findAnagram(s string, p string) []int {

	m, n := len(s), len(p)
	var arr []int
	countP := Counter(p)
	var window map[string]int

	for i := 0; i < m-n+1; i++ {
		if i == 0 {
			window = Counter(s[:n])
		} else {
			window[string(s[i-1])] -= 1
			window[string(s[i+n-1])] += 1
		}
		window = removeZero(window)
		if reflect.DeepEqual(window, countP) {
			arr = append(arr, i)
			//window = make(map[string]int)
		}
	}
	return arr
}

func main() {
	s := "cbaebabacd"
	p := "abc"
	fmt.Println(findAnagram(s, p))
}
