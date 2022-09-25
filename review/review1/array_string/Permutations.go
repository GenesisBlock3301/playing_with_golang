package main

import (
	"fmt"
	"sort"
	"strings"
)

func Sorted(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func permutationBad(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	if Sorted(s) == Sorted(t) {
		return true
	}
	return false
}

func main() {
	fmt.Println(permutationBad("nass", "sans"))
}
