// L1.8: Установка i-го бита числа int64 в 1 или 0
package main

import (
	"fmt"
)

// setBit устанавливает i-й бит числа n в значение bit (0 или 1)
func setBit(n int64, i uint, bit int) int64 {
	if bit == 1 {
		n = n | (1 << i)
	} else {
		n = n &^ (1 << i)
	}
	return n
}

func main() {
	var num int64 = 5 // 0101
	fmt.Printf("Исходное число: %d (бинарно %04b)\n", num, num)

	i := uint(1)
	result := setBit(num, i, 0)
	fmt.Printf("Устанавливаем %d бит в 0: %d (бинарная запись: %04b)\n", i, result, result)

	i = 2
	result = setBit(num, i, 1)
	fmt.Printf("Устанавливаем %d бит в 1: %d (бинарная запись: %04b)\n", i, result, result)

	i = 0
	result = setBit(num, i, 0)
	fmt.Printf("Устанавливаем %d бит в 0: %d (бинарная запись: %04b)\n", i, result, result)
}
