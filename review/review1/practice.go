package main

import "fmt"

func max(arr []int) int {
	maxVal := arr[0]
	for i := 0; i < len(arr); i++ {
		if maxVal < arr[i] {
			maxVal = arr[i]
		}
	}
	return maxVal
}

func findSmallestNumber() {
	arr := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	size := len(arr)
	output := make([]int, size)
	maxVal := max(arr) + 1
	count := make([]int, maxVal)

	//count each element
	for i := 0; i < size; i++ {
		count[arr[i]] += 1
	}

	//cumulative sum
	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}
	m := size - 1
	for m >= 0 {
		output[count[arr[m]]-1] = arr[m]
		count[arr[m]] -= 1
		m -= 1
	}
	fmt.Println(output)
}
func main() {
	arr := make([]int, 5)

	for i := 0; i < len(arr); i++ {
		arr[i] = i + 1
	}
	fmt.Println(arr[4])
	x := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(x[2:5])
	fmt.Println("Smallest number is: ")
	findSmallestNumber()

}
