package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func worker(id int, jobs <-chan string) {
	for job := range jobs {
		fmt.Printf("Worker %d got: %s\n", id, job)
	}
}

func main() {

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 0 {
		fmt.Println("Некорректное количество воркеров")
		return
	}

	data := make(chan string)

	for i := 1; i <= n; i++ {
		go worker(i, data)
	}

	i := 0
	for {
		i++
		data <- fmt.Sprintf("Сообщение %d", i)
		time.Sleep(500 * time.Millisecond)
	}
}
