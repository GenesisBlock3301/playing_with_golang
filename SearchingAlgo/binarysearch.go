package main

import (
	"fmt"
)

func binary_search(arr []int,x int) int{
	left,right := 0,len(arr)-1
	for left <= right{
		mid := (left+right)/2
		if(arr[int(mid)] == x ){
			return mid 
		}
		if arr[mid] < x{
			left = mid+1
		}else{
			right = mid -1
		}

	}
	return -1
}
func main() {
	var size int
	fmt.Printf("Enter your array size: ")
	fmt.Scanln(&size)
	var arr = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scanf("%d",&arr[i])
	}

	fmt.Printf("Enter your searching value: ")
	var val int
	fmt.Scanln((&val))
	fmt.Println(binary_search(arr,val))
}
