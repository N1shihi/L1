package main

import "fmt"

func binarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] < target {
			left = mid + 1
		} else if arr[mid] > target {
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {

	arr := []int{1, 3, 6, 7, 8, 12, 15, 16, 22, 46, 53, 64, 77}
	target := 15
	fmt.Println(binarySearch(arr, target))
	target = 14
	fmt.Println(binarySearch(arr, target))
}
