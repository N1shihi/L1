package main

import (
	"fmt"
)

func main() {
	RandWords := []string{"cat", "cat", "dog", "cat", "tree"}

	groups := make(map[string]struct{})

	for _, key := range RandWords {
		groups[key] = struct{}{}
	}

	for k := range groups {
		fmt.Println(k)
	}
}
