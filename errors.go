package main

import (
	"errors"
	"fmt"
	"os"
)

// In Go, an error is simply a value that implements the built-in error interface.
func goErrors() {
	// ex
	file, err := os.Open("order.txt")
	// Using errors.New
	err2 := errors.New("File opened")
	if err != nil {
		fmt.Println("Error occured:", err)
		return
	} else {
		fmt.Println(err2)
	}
	defer file.Close()

	// Using fmt.Errorf
	// err := fmt.Errorf("user %d not found", userID)

	// Error Wrapping
	// Error wrapping in Go is a technique that lets you add context to an existing error
	// while preserving the original error inside it.
	err3 := fmt.Errorf("failed to read file: %w", err)

	// Checking Wrapped Errors: errors.Is()
	if errors.Is(err, os.ErrNotExist) { // checks : Does err represent a ‘file does not exist’ error — even if it was wrapped
		fmt.Println("File does not exist")
	}
	// the above functio ls checks:
	// Is err equal to target?
	// OR is err wrapping something equal to target?
	// OR does anything inside the error chain match target?

	// use of errors.As()
	var myErr myError3
	if errors.As(err3, &myErr) {
		fmt.Println("Custom error with code:", myErr.code)
	}
	// the above function As does:
	// unwraps the error err3,
	// checks if it matches the type of myError3
	// if it matches it assigns to myErr

	// Sentinel Errors: A predefined error variable.
	// err4 := errors.New("Gotilla")

}

// ex
type error1 interface {
	Error() string
}

// Custom error types
type myError3 struct {
	code int
	msg  string
}

func (e myError3) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.code, e.msg)
}

// Panic vs Error (Very Important)
// 🔹 error					🔹 panic
// Expected problems				Programmer mistakes
// Should be handled				Unexpected fatal situations
// Part of normal flow

// Returning Multiple Errors
// Go doesn't support multiple errors directly
// Option 1: Return slice of errors
func validate() {
	var errors3 []error
	somethingWrong := errors.New("dont know")
	if somethingWrong == nil {
		errors3 = append(errors3, errors.New("gotilla"))
	}
}

// or use
// err := errors.Join(err1, err2)

// _ = file.Close() // intentionally ignored

// Named Error Variables in Standard Library
// Common standard errors:
// os.ErrNotExist
// io.EOF
// context.Canceled
// sql.ErrNoRows

// if errors.Is(err, io.EOF) {
//     fmt.Println("End of file")
// }

// comparing errors
// if errors.Is(err, ErrInvalidInput)
