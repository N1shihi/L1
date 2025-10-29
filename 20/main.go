package main

import (
	"fmt"
)

func reverse(b []rune, start, end int) {
	for start < end {
		b[start], b[end] = b[end], b[start]
		start++
		end--
	}
}

func reverseWords(s string) string {
	b := []rune(s)

	reverse(b, 0, len(b)-1)
	println(string(b))
	start := 0
	for i := 0; i <= len(b); i++ {
		if i == len(b) || b[i] == ' ' {
			reverse(b, start, i-1)
			start = i + 1
		}
	}

	return string(b)
}

func main() {
	input := "snow dog sun"
	output := reverseWords(input)
	fmt.Println(output)
}
