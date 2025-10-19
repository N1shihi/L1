// L1.6: Остановка горутины — все способы
package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func stopByCondition() {
	fmt.Println("1. Остановка по условию")
	done := false
	go func() {
		for {
			if done {
				fmt.Println("Горутина: завершение по условию")
				return
			}
			fmt.Println("Горутина работает...")
			time.Sleep(300 * time.Millisecond)
		}
	}()

	time.Sleep(time.Second)
	done = true
	time.Sleep(500 * time.Millisecond)
}

func stopByChannel() {
	fmt.Println("\n2. Остановка через канал")
	stop := make(chan struct{})
	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("Горутина: завершение через канал")
				return
			default:
				fmt.Println("Горутина работает...")
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()
	time.Sleep(time.Second)
	close(stop)
	time.Sleep(500 * time.Millisecond)
}

func stopByContext() {
	fmt.Println("\n3. Остановка через контекст")
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Горутина: завершение через контекст")
				return
			default:
				fmt.Println("Горутина работает...")
				time.Sleep(300 * time.Millisecond)
			}
		}
	}(ctx)

	time.Sleep(time.Second)
	cancel()
	time.Sleep(500 * time.Millisecond)
}

func stopByGoexit() {
	fmt.Println("\n4. Остановка через runtime.Goexit()")
	go func() {
		fmt.Println("Горутина начала работу")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Горутина завершает через Goexit()")
		runtime.Goexit()
		fmt.Println("Эта работа не будет выполнена")
	}()
	time.Sleep(time.Second)
}

func stopByWaitGroup() {
	fmt.Println("\n5. Ожидание завершения через sync.WaitGroup")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			fmt.Println("Горутина работает...", i)
			time.Sleep(300 * time.Millisecond)
		}
		fmt.Println("Горутина завершена")
	}()
	wg.Wait()
}

func main() {
	stopByCondition()
	stopByChannel()
	stopByContext()
	stopByGoexit()
	stopByWaitGroup()

	fmt.Println("\nВсе вариации завершены")
}
