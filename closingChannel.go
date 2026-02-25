package main

import "fmt"

// When you close a channel:
// ✅ You are telling receivers: “No more values will come.”
// ❌ You cannot send new values after closing.
// ✅ Receivers can still receive remaining buffered values.

// Important Rules of Closing Channels
// Rule 1: Only the Sender Should Close: The goroutine that sends values should close the channel.
// Rule 2: Do NOT Close from Multiple Goroutines: Closing twice causes panic
// Rule 3: Do NOT Send After Closing: causes panic
func closingChannel() {
	ch := make(chan int)
	close(ch) // closing channel

	ch2 := make(chan int)
	go func() {
		ch2 <- 12
		ch2 <- 21
		close(ch2)
	}()

	for msg := range ch2 {
		fmt.Println(msg)
	} // output 1 2 channel closed

	// What Happens When Receiving from a Closed Channel?
	// Case 1: Buffered values exist: They are received normally.
	// No values left: // Receiver gets: // Zero value of type // false status (if using comma-ok form)
	ch3 := make(chan int, 1)
	ch3 <- 100
	close(ch3)

	v1, ok1 := <-ch3
	v2, ok2 := <-ch3

	fmt.Println(v1, ok1) // 100, true
	fmt.Println(v2, ok2) // 0, false

	// Using range with Closed Channels
	for value := range ch {
		fmt.Println(value)
	}
	// This loop: // Keeps receiving // Stops when channel is closed // No panic

	// Worker pattern
	ch4 := make(chan int)
	go workerCh(ch4)
	ch4 <- 1
	ch4 <- 2

	close(ch4)

}

// Worker pattern
func workerCh(ch chan int) {
	for job := range ch {
		fmt.Println("Madathevi", job)
	}
	fmt.Println("Mugitha")
}

// | Situation           | Behavior            |
// | ------------------- | ------------------- |
// | Send after close    | ❌ Panic             |
// | Close twice         | ❌ Panic             |
// | Receive from closed | Zero value          |
// | Range over closed   | Stops automatically |
// | Close nil channel   | ❌ Panic             |
