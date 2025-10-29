package main

import (
	"fmt"
	"strings"
)

func IsUnique(str string) bool {
	str = strings.ToLower(str)
	letters := make(map[rune]bool)
	for _, r := range str {
		if letters[r] {
			return false
		}
		letters[r] = true
	}
	return true
}

func main() {
	str := "asdaf"
	fmt.Println(IsUnique(str))
}
