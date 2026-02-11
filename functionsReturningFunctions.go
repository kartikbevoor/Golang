package main

import "fmt"

// basic syntax
func outer() func() {
	return func() {
		fmt.Println("From inner function")
	}
}

// usage
// f := outer()
// f()
// or
// outer()()

// returning a function with parameters
func multiplier(factor int) func(int) int { // here multiplier is a funtion which takes the int, returns the function which inturn takes the int and returns an int
	return func(n int) int { // the inner function remembers function even after multiplier has finished execution
		return n * factor // closure concept
	}
}

// usage
// double := multiplier(2)
// triple := multiplier(3)

// fmt.Println(double(5)) // 10
// fmt.Println(triple(5)) // 15

// returning function types explicitly
type operation func(int) int

func multiplier2(factor int) operation {
	return func(n int) int {
		return factor * n
	}
}

// example
func counter() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}

// c1 := counter()
// fmt.Println(c1())  // 1
// fmt.Println(c1())  // 2

// c2 := counter()
// fmt.Println(c2())  // 1

// Returning Functions That Modify State: A function returns another function, and that returned function changes (modifies) variables it captured from its outer scope
func adder(base int) func(int) int {
	sum := base

	return func(value int) int {
		sum = sum + value
		return sum
	}
}

// usage
// a := adder(10)
// fmt.Println(a(5))  // 15
// fmt.Println(a(3))  // 18

// example
func logger(next func()) func() {
	return func() {
		fmt.Println("before execution")
		next()
		fmt.Println()
	}
}

func handler() {
	fmt.Println("Handling requests")
}

// usage
// wrapped := logger(handler)
// wrapped()

// Loop Variable Capture Problem
// for i := 0; i < 3; i++ {
//     funcs = append(funcs, func() {
//         fmt.Println(i)
//     })
// } // 3 3 3

// for i := 0; i < 3; i++ {
//     i := i // create new copy
//     funcs = append(funcs, func() {
//         fmt.Println(i)
//     })
// }  // 0 1 2

func hanga() {
	result := func(x int) func(int) int {
		return func(y int) int {
			return x + y
		}
	}(5)(10)

	fmt.Println(result) // 15

}

// Example: Retry Wrapper
func retry(times int, operation func() error) func() error {
	return func() error {
		for i := 0; i < times; i++ {
			err := operation()
			if err == nil {
				return nil
			}
		}
		return fmt.Errorf("failed after %d attempts", times)
	}
}
