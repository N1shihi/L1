package main

import (
	"fmt"
	"time"
)

func sleep(ms int) {
	ch := make(chan struct{})
	go func() {
		<-time.After(time.Duration(ms) * time.Millisecond)
		close(ch)
	}()
	<-ch
}

func main() {

	for i := 0; i < 10; i++ {
		if i == 5 {
			sleep(1000)
		}
		fmt.Printf("i=%d\n", i)
	}

}
