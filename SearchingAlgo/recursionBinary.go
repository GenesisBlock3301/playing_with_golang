package main

func binarySearch(arr []int, low int, high int, item int) int {
	var mid int
	if low <= high {
		mid = (low + high) / 2
		if arr[mid] == item {
			return mid
		} else if arr[mid] < item {
			return binarySearch(arr, mid+1, high, item)
		} else {
			return binarySearch(arr, low, mid-1, item)
		}
	} else {
		return -1
	}
}

func main() {
	arr := []int{3, 2, 1, 7}
	n := len(arr)
	println(binarySearch(arr, 0, n-1, 3))
}
