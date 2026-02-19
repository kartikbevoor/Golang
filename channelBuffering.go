package main

import "fmt"

// A channel is a typed conduit through which goroutines communicate.
func channelBuffering() {
	// ch := make(chan int)        // unbuffered channel capacity 0
	// ch := make(chan int, 3)     // buffered channel (capacity 3) this means the channel can store upto 3 values.

	// Unbuffered Channel: Send blocks until another goroutine receives.,
	// Receive blocks until another goroutine sends.,
	// Sender and receiver must be ready at the same time.

	ch := make(chan int)

	go func() {
		ch <- 10 // someone is sending: blocks until someone receives
	}()

	value := <-ch // someone is receiving
	fmt.Println(value)
	// This creates synchronization between goroutines.
	// No value is stored — it is handed off directly.

	chUnbuffered := make(chan int, 2)
	// Send blocks only when buffer is full and Receive blocks only when buffer is empty

	chUnbuffered <- 1
	chUnbuffered <- 2
	fmt.Println("done")

	// now we cant do
	//chUnbuffered <- 3 // becos channel is full, we need a receiver.

	fmt.Println(len(chUnbuffered)) // 2 (current elements)
	fmt.Println(cap(chUnbuffered)) // 3 (total capacity)

	// Unbuffered channel → Direct phone call
	// Buffered channel → Voice mailbox with limited storage
}
