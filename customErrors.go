package main

import (
	"errors"
	"fmt"
)

// In Go, errors are values, not exceptions.
func customErrors() {

	// creating basic custom errors
	// using errors.New()  var ErrNotFound = errors.New("not found")
	// fmt.Errorf()		   err := fmt.Errorf("user %d not found", id)

	// type assertion with custom errors.
	err := validateUser(15)

	if err != nil {
		if vErr, ok := err.(validationError); ok { // This is called a type assertion. It asks, Is the interface value err actually holding a value of type ValidationError
			fmt.Println("Field:", vErr.field)
			fmt.Println("Message:", vErr.msg)
		}
	}

	// Error chains when wrapping multiple times
	err1 := errors.New("New error")
	err2 := fmt.Errorf("Wrapped error1 : %w", err1)
	err3 := fmt.Errorf("Wrapped error2: %w", err2)
	fmt.Println(errors.Is(err1, err3)) // false
	fmt.Println(errors.Is(err3, err1)) // true

	// Multiple error wrapping
	err4 := errors.Join(err1, err2)
	fmt.Println(errors.Is(err4, err2))

}

// The built-in error type is an interface:
// type error interface {
//     Error() string
// }
// so any type that implements Error() string is considered an error

// custom errors using struct
type validationError struct {
	field string
	msg   string
}

func (ve validationError) Error() string {
	return fmt.Sprintf("Validation failed on %s: %s", ve.field, ve.msg)
}

func validateUser(age int) error {
	if age < 18 {
		return validationError{
			field: "age",
			msg:   "user not allowed",
		}
	}
	return nil
}

// Using Pointer vs Value Receiver
// func (e *ValidationError) Error() string {
//     return ...
// }

// The Idea Behind Wrapping
// Sometimes you want to:
// Add context to an error
// But NOT lose the original error

// Custom Error with Unwrap()
// func (e *myError3) Unwrap() error {
//     return e.msg
// }

type AppError struct {
	Message string
	Err     error
}

func (e *AppError) Error() string {
	return e.Message + ": " + e.Err.Error()
}

func (e *AppError) Unwrap() error { // Unwrap() by the adding this: we are telling go,
	return e.Err // This error wraps another error, If someone wants the inner error, return it.
}

// ex
type notEnoughBalance struct {
	balance int
	amount  int
}

func (ne *notEnoughBalance) Error() string {
	return fmt.Sprintf(
		"Insufficient balance: %d, amount tried to withdraw: %d",
		ne.balance, ne.amount)
}

func checkBalance(amt int, bal int) error {
	if amt > bal {
		return &notEnoughBalance{
			balance: bal,
			amount:  amt,
		}
	}
	return nil
}
