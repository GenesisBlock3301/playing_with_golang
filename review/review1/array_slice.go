package main

import "fmt"

func Array() {
	x := [5]int64{1, 2, 3, 4, 5}
	fmt.Println(x)
}

func sliceArray() []float64 {
	//var x[]float64
	x := make([]float64, 5, 10)
	return x
}

func append_copy() []int {
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	return slice2
}

func appendUserInput() []int {
	var n int
	fmt.Println("Enter your Number: ")
	fmt.Scanf("%d", &n)
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Enter %dth element: ", i)
		fmt.Scanf("%d", &arr[i])
	}

	return arr
}

func copy_array() ([]int, []int) {
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 5, 6)
	copy(slice2, slice1)
	return slice2, slice1
}

func main() {
	Array()
	fmt.Println(sliceArray())
	fmt.Println(append_copy())
	//fmt.Println("user inputed array: ")
	//fmt.Println(appendUserInput())
	fmt.Println(copy_array())
}
