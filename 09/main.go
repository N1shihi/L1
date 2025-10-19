package main

import (
	"fmt"
)

func stage1(nums []int, out chan<- int) {
	for _, x := range nums {
		out <- x
	}
	close(out)
}

func stage2(in <-chan int, out chan<- int) {
	for x := range in {
		out <- x * 2
	}
	close(out)
}

func main() {
	nums := []int{1, 2, 3, 4, 5}

	ch1 := make(chan int)
	ch2 := make(chan int)

	go stage1(nums, ch1)

	go stage2(ch1, ch2)

	for result := range ch2 {
		fmt.Println(result)
	}

	fmt.Println("Конвейер завершён корректно.")
}
