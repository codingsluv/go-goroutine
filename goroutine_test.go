package belajargoroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func DisplayNumber(number int) {
	fmt.Println("Display :", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go DisplayNumber(i)
	}
	time.Sleep(1 * time.Second) // wait for goroutines to finish
	t.Log("All goroutines finished")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("ups")
	time.Sleep(1 * time.Second)
}
