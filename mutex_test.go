package belajargoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	counter := 0
	var mutex sync.Mutex

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				counter = counter + 1
				mutex.Unlock()
			}
		}()
	}
	time.Sleep(3 * time.Second)
	fmt.Println("Counter =", counter)
}
