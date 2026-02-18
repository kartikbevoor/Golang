package main

import "fmt"

// Desclaimer: this file only covers basics of channels cover it later
func channels() {
	// channel declaration
	// var ch chan int

	// channel creation
	ch := make(chan int)

	ch <- 42 // puts 42 into the channel

	value := <-ch
	fmt.Println(value)

	// unbuffered channels
	// ch := make(chan int)

	// buffered channels
	// ch := make(chan int, 3)	// buffer size 3

}

// simple communication
func simpleCommunication() {
	ch := make(chan string)

	go func() {
		ch <- "Hello"
	}()

	msg := <-ch
	fmt.Println(msg)
}

// Directional Channels: Restrict usage to send-only or receive-only.
func sendData(ch chan<- int) {
	ch <- 10 // sending data into channel
}

func reciveData(ch <-chan int) int {
	return <-ch
}

// goroutines communicate via channels

//   +-------------------+                 +-------------------+
//   |   Goroutine A      |                 |   Goroutine B      |
//   | (Producer/Sender)  |                 | (Consumer/Receiver)|
//   +-------------------+                 +-------------------+
//              |                                   ^
//              | sends data via channel (ch)      |
//              |--------------------------------->|
//              |                                   |
//              |                                   |
//              v                                   |
//       [   Channel ch (buffered/unbuffered)    ]

// 	Explanation:
// Goroutine A produces data and sends it into the channel using ch <- data.
// Channel ch acts as a pipe.
// If unbuffered, A waits until B receives the data.
// If buffered, A can send up to the buffer capacity without waiting.
// Goroutine B receives data using data := <-ch.
// Communication is synchronized, so no explicit locks are needed.
