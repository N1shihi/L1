package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Укажите количество секунд, пример: go run main.go <количество_секунд>")
		return
	}

	N, err := strconv.Atoi(os.Args[1])
	if err != nil || N <= 0 {
		fmt.Println("Некорректное количество секунд")
		return
	}

	data := make(chan int)

	go func() {
		for v := range data {
			fmt.Printf("Получено: %d\n", v)
		}
		fmt.Println("Чтение завершено")
	}()

	timeout := time.After(time.Duration(N) * time.Second)

	i := 0

loop:
	for {
		select {
		case <-timeout:
			fmt.Println("Время истекло")
			break loop
		default:
			i++
			data <- i
			time.Sleep(500 * time.Millisecond)
		}
	}

	close(data)
}
