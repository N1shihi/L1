package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	for i := len(arr) - 1; i >= 0; i-- {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			fmt.Printf("%d^2 = %d\n", num, num*num)
		}(arr[i])
	}
	wg.Wait()
}
