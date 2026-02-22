package main

import (
	"fmt"
)

// Panic is a built-in function that stops normal execution of the current function
//  and begins unwinding the stack.

// It is used for serious runtime errors that the program cannot recover from normally.
// Something went terribly wrong — stop everything!

// panic is used when: The program reaches an unexpected state, There is a critical programming error
// Execution cannot continue safely.

// When panic is called: Current function stops immediately, Deferred functions run,
// Stack unwinds (caller functions exit one by one), Program crashes (unless recovered)

// Automatically (Runtime Panic): Go automatically panics in situations like:
// Index out of range, Nil pointer dereference, Divide by zero

// Manually using panic function

// If panic occurs: It stops current function, Goes up the call stack,
// Executes all deferred functions, Crashes at top level if not recovered
func panicGo() {
	fmt.Println("Start")
	panic("Something went wrong!") // Manually using panic function
	// fmt.Println("End") // This will NOT execute

	// index out of bound panic
	// arr1 := [3]int{1, 2, 3}
	// fmt.Println(arr1[4])

	// test()

	// Panic teversal
	// aPanic()
	// Program begins execution in this order
	// aPanic() -> bPanic() ->panic()
	// Apfter panic
	// B → A → main → program crash

}

func test() {
	defer fmt.Println("defer 1")
	defer fmt.Println("defer2")
	panic("Error")
} // defer 2, defer1, Error (execution)
// Note : After executing defers, program crashes

// | panic                         | error                      |
// | ----------------------------- | -------------------------- |
// | Crashes program               | Does NOT crash program     |
// | Used for unrecoverable issues | Used for expected failures |
// | Rarely used in normal code    | Very commonly used         |
// | Stops execution immediately   | Returned as value          |

func dividePanic(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

// Recovering from panic
// Note: Recover only works in deferred function
func recoverFromPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from:", r)
		}

		// recover() this function stops the panic and returns the value passed to panic

		// The above function is same as below
		// r := recover()
		// if r != nil {
		// 	fmt.Println("Recovered from:", r)
		// }
	}()

	panic("Something went wrong")
} // Recovered from: Something went wrong

// Panic traversal:
func aPanic() {
	bPanic()
}

func bPanic() {
	panic("Panic")
}

// ex
func mustConnect(connectionFailed bool) {
	if connectionFailed {
		panic("Database connection failed")
	}
}

// panic() called
//        ↓
// Stop normal execution
//        ↓
// Run deferred functions (LIFO)
//        ↓
// If recover() exists → resume execution
// Else → crash program
