package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"
)

// Используем контекст, потому что этот способ идиоматичен, он сочетается с библиотеками, поддерживающими контекст
// Принцип работы: при получении SIGINT вызывается cancel, main пеерестает генерировать задания и закрывает канал jobs
func worker(id int, jobs <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for msg := range jobs {
		fmt.Printf("Worker %d got: %s\n", id, msg)
	}
	fmt.Printf("Worker %d exiting\n", id)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Укажите количество воркеров, пример: go run main.go <количество_воркеров>")
		return
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 0 {
		fmt.Println("Некорректное количество воркеров")
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	data := make(chan string, 100)

	var wg sync.WaitGroup

	for i := 1; i <= n; i++ {
		wg.Add(1)
		go worker(i, data, &wg)
	}

	go func() {
		<-sigCh
		fmt.Println("\nSIGINT received — shutting down...")
		cancel()
		signal.Stop(sigCh)
	}()

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	msgID := 0
loop:
	for {
		select {
		case <-ctx.Done():
			break loop
		case <-ticker.C:
			msgID++
			data <- fmt.Sprintf("Сообщение %d", msgID)
		}
	}

	close(data)

	wg.Wait()

	fmt.Println("All workers exited, program terminating.")
}
