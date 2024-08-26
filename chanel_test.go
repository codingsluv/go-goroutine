package belajargoroutine

import (
	"fmt"
	"strconv"
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

func TestBufferdChannel(t *testing.T) {
	ch := make(chan string, 2)
	defer close(ch)

	go func() {
		ch <- "hello"
		ch <- "world"
	}()

	go func() {
		fmt.Println(<-ch)
		fmt.Println(<-ch)
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("Selesai")
	t.Log("All goroutines finished")
}

func TestRangeChannel(t *testing.T) {
	ch := make(chan string)

	go func() {
		for i := 0; i < 11; i++ {
			ch <- "looping ke :" + strconv.Itoa(i)
		}
		close(ch)
	}()

	for data := range ch {
		fmt.Println(data)
	}
	fmt.Println("done")
}

func TestSelectChannel(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)
	defer close(ch1)
	defer close(ch2)

	go GiveMeResponse(ch1)
	go GiveMeResponse(ch2)

	counter := 0
	for {
		select {
		case data := <-ch1:
			fmt.Println("data dari channel 1", data)
			counter++
		case data := <-ch2:
			fmt.Println("data dari channel 2", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}
}
