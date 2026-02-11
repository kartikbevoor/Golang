package main

// continue from 19

import (
	"errors"
	"fmt"
)

// simple function
func add(a int, b int) {
	fmt.Println(a + b)
}

// function with return satement
func add2(a int, b int) int {
	return a + b
}

// multiple parameters with same type
func add3(a, b int) int {
	return a + b
}

// function with no return statement and parameters
func sayHello() {
	fmt.Println("Hello")
}

// function with multiple return value
func divide(a, b int) (int, int) {
	quotient := a / b
	reminder := a % b
	return quotient, reminder
}

// q, r := divide(10, 3) this is how the above function is called
// q, _ := divide(10, 3) ignoring the return value

// function with named return value
func areaRectangle(length int, width int) (area int) {
	area = length * width
	return // implicit return
}

func isPositive(n int) bool {
	if n <= 0 {
		return false
	} else {
		return true
	}
}

// function returning errors
func findQuotient(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Division not possible by zero") // make sure the order must follow here int after error
	}
	return a / b, nil
}

// Above code's usage
// result, err := divideSafe(10, 0)
// if err != nil {
//     fmt.Println("Error:", err)
//     return
// }
// fmt.Println(result)

// Variadic function(variable number of arguments)
func add4(numbers ...int) int {
	total := 0
	for _, v := range numbers {
		total = total + v
	}
	return total
}

// calls for the above function
// sum(1, 2, 3)
// sum(1, 2, 3, 4, 5)

// passing a slice
// nums := []int{1, 2, 3}
// sum(nums...)

// function as values
// func add(a, b int) int {
//     return a + b
// }

// var operation func(int, int) int
// operation = add

// fmt.Println(operation(2, 3)) // 5

// defered function calls
func example() { // look deep
	defer fmt.Println("hello") // second this statement runs
	fmt.Println("world")       // first this statement runs
}

// note: Multiple defers (LIFO order)
//example
// defer fmt.Println(1)
// defer fmt.Println(2)
// defer fmt.Println(3)
// 3 2 1

// Methods: when a function is attached to struct it becomes method
type person struct {
	name string
	age  int
}

func (p person) methods() {
	fmt.Println(p.name, p.age)
}

// usage
// p := Person{name: "sam", age: 23}
// p.methods()

func (p person) changeName() {
	p.name = "fam"
}

// The above function does not modify the original

func (p *person) changeName2() {
	p.name = "fam"
}

// the above function modifys the original

// Init function: runs automatically, runs before main, cannot be called explicitly
func init() {
	fmt.Println("Init function which runs automatically, runs before main and cannot be called explicitly")
}

// note: we can have multiple init functions

// note: functions whose name is stated using capital letters those are called exported functions and they are public functions
