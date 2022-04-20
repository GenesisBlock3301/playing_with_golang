package main

import (
	"fmt"
)
func linearSearch(arr []int,x int) (int){
	n := len(arr)
	i := 0
	for i < n {
		if(arr[i] == x){
			return i
		}
		i++
	}
	i = -1
	return i
}
func main(){
	fmt.Printf("Enter size of array: ")
	var size int
	fmt.Scanln(&size)
	var arr = make([]int,size)
	for i:=0;i<size;i++{
		fmt.Printf("Enter %dth element ",i)
		fmt.Scanf("%d",&arr[i])
	}
	fmt.Println()
	fmt.Printf("Enter your searching value: ")
	var x int
	fmt.Scanln(&x)
	fmt.Println(linearSearch(arr,x))
}