package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// Timeouts in Go are used to limit how long an operation is allowed to run.
func timeouts() {

	// Timeout Using time.After(duration)

	// Timeout Using context.WithTimeout()
	// ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

}

// Using context.WithTimeout
func contextWithTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	ch := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- "Finished"
	}()

	select {
	case msg := <-ch:
		fmt.Println(msg)
	case <-ctx.Done():
		fmt.Println("Timeout:", ctx.Err())
	}
}

// context.Background(): This creates a root context - The starting point for all contexts.
//   Never cancels, Has no timeout, Has no values
// context.WithTimeout(...): This function: Takes a parent context
//   Returns: A new derived context, A cancel function

// The whole line means: Create a new context that will automatically cancel after 2 seconds.

// ctx → of type context.Context
// cancel → of type context.CancelFunc (which is func())
// context.Background
// It returns a base context:
// It is never canceled
// It has no timeout
// It has no values
// It is typically used at the top level (like in main())
// Think of it as the root of a context tree
// context.WithTimeout(parentContext, duration)
// Creates a new child context
// That automatically cancels after the given duration
// Returns:
// The new context
// A cancel function
// Internally: func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
// Creates a new context struct
// Links it to the parent (Background)
// Starts a timer for 2 seconds
// When the timer fires:
// It cancels the context
// Closes the Done() channel
// Sets the error to context.DeadlineExceeded

// Timeout for http clients:
func timeoutForHttpClient() {
	client := http.Client{
		Timeout: 2 * time.Second,
	}
	fmt.Println(client)

}

// resp, err := client.Get("https://example.com") // if this request exceeds 2 secs error occurs

// Timeout Using time.NewTimer()
func timeoutUsingtimeNewTimer() {
	timer := time.NewTimer(2 * time.Second)
	defer timer.Stop()

	select {
	case <-timer.C:
		fmt.Println("Timeout")
	case <-timer.C:
		fmt.Println("Timeout")
	}

}

// Timeouts = Defensive Programming in Concurrency
