package main

import (
	"fmt"
)

func intersection(a, b []int) []int {
	m := make(map[int]struct{})
	for _, val := range a {
		m[val] = struct{}{}
	}

	var result []int
	for _, val := range b {
		if _, exists := m[val]; exists {
			result = append(result, val)
		}
	}

	return result
}

func main() {
	A := []int{1, 2, 3}
	B := []int{2, 3, 4}

	res := intersection(A, B)
	fmt.Printf("A = %v\nB = %v\nПересечение = %v\n", A, B, res)
}
