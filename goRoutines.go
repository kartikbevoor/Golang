package main

import (
	"fmt"
	"sync"
	"time"
)

// A goroutine is a lightweight thread managed by the Go runtime.
// It runs functions concurrently.
// Much cheaper than OS threads.
// Thousands or even millions can run simultaneously.
// In simple words: “A function running independently alongside other functions.”
// a thread is the smallest unit of execution within a program.
//	Thread: It represents a single sequence of instructions that the processor can execute.

// goroutines:
// Start with small stack (2KB)
// Stack grows/shrinks dynamically
// Multiplexed onto OS threads
func goRoutines() {

	// syntax:
	// go functionName()  // use of go keyword
	go sayHello()
	time.Sleep(time.Second) // Wait so goroutine can finish
	// Note: the above wait is due to, the go routines after the function in which they are called
	// suppose if goroutine is written in main funtion and its last there may pasibility that it may never run

	// Anonymous Goroutines
	go func() {
		fmt.Println("go routine")
	}()
	time.Sleep(time.Second)

	// execution order is not guarenteed in go
	for i := 0; i < 4; i++ {
		go fmt.Println(i)
	} // execution could be any order not predictable

	// Loop variable problem
	// all go routines may print the same value becos they capture the same value
	for _, v := range []int{1, 2, 3} {
		go func() {
			fmt.Println(v)
		}()
	}

	// corrected for the above
	for _, v := range []int{1, 2, 3} {
		v := v
		go func() {
			fmt.Println(v)
		}()
	} // or just pass v as parameter

	// Synchronization (Waiting for Goroutines)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println("Go")
	}()

	wg.Wait()
	// Add(n) → number of goroutines
	// Done() → decreases counter
	// Wait() → blocks until counter = 0

	// Race Condition : Multiple goroutines accesing the same shared data atleast one of them
	// is trying to modify that data
	var count int

	go func() {
		count++
	}()

	// Using mutex to prevent race condition
	var mu sync.Mutex
	var count2 int

	go func() {
		mu.Lock()
		count2++
		mu.Unlock()
	}()

}

// ex
func sayHello1() {
	fmt.Println("Hello")
}

// Go Scheduler (How Goroutines Run):
// Go uses a M:N scheduler:
// Many Goroutines (G)
// Mapped onto fewer OS threads (M)
// Controlled by logical processors (P)

// The scheduler:
// Distributes goroutines
// Preempts long-running ones
// Balances workload

// comunication between goroutines
func comGo() {
	ch := make(chan int)

	go func() {
		ch <- 23
	}()

	value := <-ch
	fmt.Println(value)
}

// Channels:
// Prevent race conditions
// Enable safe communication
// Encourage "share memory by communicating"

// Concurrency: Multiple tasks in progress (may not run simultaneously).
// Parallelism: Tasks actually run at the same time (multi-core CPU).

// Parallelism can be controled using: runtime.GOMAXPROCS(n)
// runtime.GOMAXPROCS(n) is a function call in Go (Golang) that sets the
// maximum number of OS threads that can execute Go code simultaneously

// Goroutine Leaks: happens when a goroutine is started but never finishes
// it stays blocked forever, consuming memory and scheduler resources.
// Reasons for goroutines leak
// Goroutine Waits Forever
func worker(ch chan int) {
	fmt.Println("Waiting for data")
	value := <-ch
	fmt.Println(value)
}
func alternateMain() {
	ch := make(chan int)
	go worker(ch) // no data in ch so no data sent to worker function, so goroutine waits forever.
}

func noReceiver() {
	ch := make(chan int)

	go func() {
		ch <- 10 // blocks forever
	}()

	// no receiver
}

// No Exit Condition (Infinite Loop)
func noExitCondition() { // thsi function runs forever
	for {
		time.Sleep(time.Second)
	}
}
