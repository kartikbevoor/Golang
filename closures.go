package main

import (
	"fmt"
	"net/http"
)

// Closure: A function value that references variables from outside its own function body.
// A closure = function + captured variables
// It “remembers” the variables from the scope where it was created.
// A closure is a function that captures and remembers variables from its surrounding scope.
func closures() {
	x := 10

	add := func(y int) int { // ananomous funtion assigned to a variable
		return x + y
	}
	// here add is a ananomous function, which uses x outside its scope

	// Closures capture variables(not value)
	y := 10
	fmt.Println(add(5)) // output: 15

	f := func() {
		fmt.Println(y)
	}

	y = 25
	f()
}

// Returning closures
func counter2() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}

// usage
// c := counter()
// fmt.Println(c()) // 1
// fmt.Println(c()) // 2
// fmt.Println(c()) // 3
// when each call modifies same variable it is called stateful closure; in above case

// Each closure has its own state
// c1 := counter()
// c2 := counter()
// fmt.Println(c1()) // 1
// fmt.Println(c1()) // 2
// fmt.Println(c2()) // 1

// Closures in Loops
func hanga2() {
	nums := []int{10, 20, 30}

	for _, v := range nums {
		// v := v // this line fixes the issue
		go func() {
			fmt.Println(v)
		}() // these backet is nothing but a function call
	}
	// expected output: 10 20 30 actual output: 30 30 30
	// this is becos closures capture variables not value and goroutines run after the loop
	// by the time loop finish execution v = 30
}

// example: Closures Are First-Class Values
func apply(f func(int) int, value int) int {
	return f(value)
}

func hange() {
	double := func(x int) int {
		return x * 2
	}

	fmt.Println(apply(double, 5)) // 10
}

// http handlers
func handler1(name string) http.HandlerFunc { // is the return fuction, http.HandlerFunc is a function type, that looks like func(http.ResponseWriter, *http.Request)
	return func(w http.ResponseWriter, r *http.Request) { // this function is returned
		fmt.Fprintf(w, "Hello %s", name)
	}
}

// Middleware is: A function that wraps another HTTP handler to add extra behavior.
// It does something before and/or after calling the next handler.
func logger1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request received")
		next.ServeHTTP(w, r)
	})
}
