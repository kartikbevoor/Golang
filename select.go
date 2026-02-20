package main

import (
	"fmt"
	"time"
	// "time"
)

// Select is used to wait on multiple channel operations.
// It allows a goroutine to wait until one of multiple channels is ready,
// making it essential for concurrency and coordination.
// think of select as a switch statement — but for channels.
func selectGo() {
	//sel()
	//receiveMulChan()
}

// Receiving from multiple channels
func receiveMulChan() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		ch1 <- "From 1"
	}()

	go func() {
		time.Sleep(time.Second)
		ch2 <- "From 2"
	}()

	select {
	case msg := <-ch1:
		fmt.Println(msg)
	// case <-time.After(2 * time.Second):
	// 	fmt.Println("time")
	case msg := <-ch2:
		fmt.Println(msg)
		// default: // dont use this, this is ready before everyone
		// 	fmt.Println("No msg from any one")
	}
}

// ex: 1
func sel() {
	// basic syntax
	ch := make(chan string, 1)
	// ch <- "select"
	go func() {
		ch <- "I m here"
	}()

	go func() {
		fmt.Println(<-ch)
	}()

	// time.Sleep(time.Second) // I m here, sent both are of diff case

	select {
	case ch <- "hello":
		fmt.Println("sent")
	case msg := <-ch:
		fmt.Print("Received msg: ")
		fmt.Println(msg)
	default:
		fmt.Println("No channel ready")
	}
	// time.Sleep(time.Second) // sent, hello these are of case one
}

// Using select inside loop
func selectInsideLoop() {
	ch := make(chan string)

	for {
		select {
		case msg := <-ch:
			fmt.Println(msg)
		case <-time.After(2 * time.Second):
			fmt.Println("Time over")
			return
		}
	}
}

// Handling channel closure
func handlingChannelClosure(ch chan int) {
	select {
	case msg, ok := <-ch:
		time.Sleep(2 * time.Second)
		if !ok {
			fmt.Println("channel closed")
			return
		} else {
			fmt.Println(msg)
		}
	default:
		fmt.Println(" ")

	}
} // simplest one is written in channels.go file

// Timeout pattern
func timeout(ch chan int) {
	select {
	case msg := <-ch:
		fmt.Println(msg)
	case <-time.After(2 * time.Second):
		fmt.Println("Mugit nadi")
	}
}

// Multiple ready channels
func multiReadyChannel(ch1 chan int, ch2 chan int) {
	select {
	case <-ch1:
		fmt.Println(ch1)
	case <-ch2:
		fmt.Println(ch2)
	}
} // Go randomly chooses one., This prevents starvation and ensures fairness.

// Rule      					-			Explanation
// Only works with channels					Cannot use normal variables
// Blocks by default						Unless default is present
// Random case selection					If multiple ready
// Nil channel blocks forever				Case with nil channel is ignored(will never execute, even if its the only one)
// Empty select blocks forever				select {} blocks forever

// Worker Listening to Multiple Signals
func workerListiningMultipleSignals(jobs chan int, quit chan int) {
	for {
		select {
		case job := <-jobs:
			// process(job)
			fmt.Println(job)

		case <-quit:
			fmt.Println("Stopping worker")
			return
		}
	}
}

// Why select is Important
// Without select, you cannot:
// Listen to multiple channels simultaneously
// Implement timeouts
// Create event loops
// Build scalable concurrent systems
// Handle cancellation signals
