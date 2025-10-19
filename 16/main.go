package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[len(arr)/2]
	var left, right []int

	for _, v := range arr {
		if v < pivot {
			left = append(left, v)
		} else if v > pivot {
			right = append(right, v)
		}
	}

	return append(append(quickSort(left), pivot), quickSort(right)...)
}

func main() {
	arr := []int{10, 5, 2, 3, 7, 6, 8, 1, 9, 4}
	fmt.Println("Исходный массив:", arr)

	sorted := quickSort(arr)
	fmt.Println("Отсортированный массив:", sorted)
}
