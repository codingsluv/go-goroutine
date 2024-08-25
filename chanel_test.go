package belajargoroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	ch := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		// ? kirim data ke channel
		ch <- "Hello World"
		fmt.Println("selesai kirim data ke channel")
	}()

	// ? ambil data dari channel
	data := <-ch
	fmt.Println(data)
	defer close(ch)
}

func GiveMeResponse(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "data di dalam channel"
}

func TestChannelAsParameter(t *testing.T) {
	ch := make(chan string)
	defer close(ch)

	go GiveMeResponse(ch)

	data := <-ch
	fmt.Println(data)
	time.Sleep(1 * time.Second) // wait for goroutine to finish
	t.Log("All goroutines finished")
}

// ? Mengirim data ke channel
func OnlyIn(ch chan<- string) {
	ch <- "Hello World"
	fmt.Println("selesai kirim data ke channel") // ? menampilkan data
}

// ? Menerima data dari channel
func OnlyOut(ch <-chan string) {
	data := <-ch
	fmt.Println(data) //? menampilkan data
}

func TestInOutChannel(t *testing.T) {
	ch := make(chan string)
	defer close(ch)

	go OnlyIn(ch)
	go OnlyOut(ch)

	time.Sleep(1 * time.Second) // wait for goroutines to finish
	t.Log("All goroutines finished send in out")

}
