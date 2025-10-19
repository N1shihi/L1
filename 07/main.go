// L1.7: Безопасная запись в map (конкурентный доступ)
package main

import (
	"fmt"
	"sync"
	"time"
)

// 1. Вариант с sync.Mutex
type SafeMap struct {
	mu sync.Mutex
	m  map[string]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{m: make(map[string]int)}
}

func (s *SafeMap) Set(key string, value int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key] = value
}

func (s *SafeMap) Get(key string) (int, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	val, ok := s.m[key]
	return val, ok
}

func main() {
	fmt.Println("=== Вариант с sync.Mutex ===")
	safeMap := NewSafeMap()
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			safeMap.Set(fmt.Sprintf("key-%d", i), i)
			val, _ := safeMap.Get(fmt.Sprintf("key-%d", i))
			fmt.Printf("key-%d => %v\n", i, val)
		}(i)
	}

	wg.Wait()

	fmt.Println("Все записи завершены без гонок (проверьте: go run -race main.go)")

	time.Sleep(200 * time.Millisecond)
	fmt.Println("\nПрограмма завершена.")
}
