package main

import "fmt"

func maxVal(arr []int) int {
	mx := arr[0]
	for _, val := range arr {
		if val > mx {
			mx = val
		}
	}
	return mx
}

func valExistence(m int, arr []int) bool {
	for _, val := range arr {
		if m == val {
			return true
		}
	}
	return false
}

func isValid(matrix [][]int) bool {
	rows := make([][]int, len(matrix))
	cols := make([][]int, len(matrix))
	mx_n := maxVal(matrix[0])
	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix); c++ {
			m := matrix[r][c]
			if 1 <= m && m <= mx_n {
				//	check existence
				if valExistence(m, rows[r]) || valExistence(m, cols[c]) {
					return false
				}
			}
			if len(rows) != len(cols) {
				return false
			}
			rows[r] = append(rows[r], m)
			cols[c] = append(cols[c], m)
		}
	}

	return true
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{3, 1, 2},
		{2, 3, 1},
	}

	fmt.Println(isValid(matrix))
}
