package main

import "fmt"

func main() {

	nums := []int{10, 20, 30, 40, 50}
	del := 2

	fmt.Println("До удаления:", nums)

	copy(nums[del:], nums[del+1:])

	nums = nums[:len(nums)-1]

	fmt.Println("После удаления:", nums)
}
