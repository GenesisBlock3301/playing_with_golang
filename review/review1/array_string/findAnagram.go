package main

import (
	"fmt"
	"sort"
	"strings"
)

func isElementExist(str string, counter map[string][]string) bool {
	for key, _ := range counter {
		if str == key {
			//fmt.Println(str, key, counter[key])
			return true
		}
	}
	return false
}

func sortingKey(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func mapToArray(hash_map map[string][]string) [][]string {
	var arr [][]string
	for _, val := range hash_map {
		arr = append(arr, val)
	}
	return arr
}
func findAnagram1(strs []string) [][]string {
	hash_map := make(map[string][]string)

	for _, val := range strs {
		if isElementExist(val, hash_map) {
			s := sortingKey(val)
			fmt.Println(s, val)
			hash_map[s] = append(hash_map[s], val)
			//	fmt.Println(val)
		} else {
			s := sortingKey(val)
			hash_map[s] = append(hash_map[s], val)
		}
	}
	return mapToArray(hash_map)
}
func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	//test_str := map[string][]string{
	//	"nas": []string{"hello", "sifat"},
	//}
	//fmt.Println(isElementExist("nas", test_str))
	fmt.Println(findAnagram1(strs))

}
