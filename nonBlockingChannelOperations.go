package main

import "fmt"

// In Go, channel operations are blocking by default:
// Sending to a channel blocks until a receiver is ready.
// Receiving from a channel blocks until a sender sends something.

// Non-blocking operations let you:
// Try sending without waiting
// Try receiving without waiting
// Continue execution if channel isn’t ready
func nonBlockingChannelOperation() {
	ch := make(chan int)
	// ch <- 12 // blocks forever no receiver If no goroutine is receiving, the program deadlocks.

	// using select + default for non blocking
	select {
	case ch <- 13:
		// sent
	default:
		// unable to send
	} // If the send/receive cannot proceed immediately, default executes.

	// Non-Blocking Receive
	select {
	case msg := <-ch:
		fmt.Println(msg)
	default:
		fmt.Println("No msg recieved")
	} // here default block is executed

	// Non-Blocking Send
	select {
	case ch <- 21:
		fmt.Println("sent")
	default:
		fmt.Println("agalilla")
	} // here default block is executed

	// In buffered channels
	ch2 := make(chan int, 2)
	select {
	case ch2 <- 21:
		fmt.Println("Sent successfully")
	default:
		fmt.Println("Channel is fulled")
	} // default is executed when channel is fulled

}
