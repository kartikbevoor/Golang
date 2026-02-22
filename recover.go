package main

import (
	"fmt"
	"net/http"
)

// Recover is a built-in function in Go used to regain control of a goroutine after a panic.
// It stops a panic and prevents the program from crashing

// recover() allows you to: Catch the panic, Handle it gracefully, Continue program execution
// recover(): Returns the value passed to panic, Returns nil if there is no active panic

// Note: recover() works ONLY inside a deferred function, If called outside defer, it does nothing.
// recover outside defer does nothing
func recoverGo() {
	go recoverGoroutines() // Each goroutine must recover its own panic.
	// If you don't recover inside the goroutine → entire program crashes.
}

// ex panic with recover
func panicWithRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from:", r)
		}
	}()
	fmt.Println("Start")
	panic("Something went wrong")
	// fmt.Println("End") // never runs
}

// recover outside defer does nothing
func recoverOutsideDefer() {
	recover()
	panic("what is wrong with you")
}

// Recover with goroutines
func recoverGoroutines() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from:", r)
		}
	}()
	panic("Wrong")
}

// ex: Web Servers (like HTTP middleware)
func handlerRecover(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			http.Error(w, "Internal Server Error", 500)
		}
	}()
}
