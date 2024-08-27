package belajargoroutine

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println(message)
	}
}

func TestPrint(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	go print(5, "Hello, World!")
	go print(5, "Belajar Go Routine!")
	time.Sleep(1 * time.Second)
	fmt.Println("Finished")
}

func TestGetGomaxProcs(t *testing.T) {
	totalCpu := runtime.NumCPU()
	fmt.Printf("Total CPU: %d\n", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Printf("Total Thread: %d\n", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Printf("Total Goroutine: %d\n", totalGoroutine)
}
