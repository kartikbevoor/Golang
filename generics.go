package main

import (
	"fmt"
)

// Generics allow you to write functions, types, and data structures that work with multiple types
// while preserving type safety.
func generics() {
	fmt.Println(addGen(2, 3.5))
	fmt.Println(Identity("jam"))
	fmt.Println(Identity(99))

	// comparable constraint
	fmt.Println(compare("aaa", "aba"))
	fmt.Println(compare(10, 10))

	// Generic structs
	b1 := box[int]{
		value: 90,
	}
	fmt.Println(b1.value)

	b2 := box[string]{
		value: "yen ri media",
	}
	fmt.Println(b2.value)

	// generic methods
	b1.returnValue()

	// Explicit usage:
	// addGen[int](3, 4)

	// generic stack
	s := stack[int]{}
	s.push(10)
	s.push(20)

	// Generic map function
	// nums := []int{1, 2, 3}
	// squares := Map(nums, func(x int) int {
	// 	return x * x
	// })
	// the above is error check

}

// ex
func addGen[t int | float64](a t, b t) t { // Type Constraints : Constraints restrict what types are allowed. // only allows int and float64
	return a + b
}

// syntax
// func FunctionName[T constraint](param T) T

func Identity[T any](value T) T { // any keyword : Works with any type.
	return value
}

// Union Type Constraints
type allowedNum interface {
	int | float64 | float32
}

func subGen[T allowedNum](a, b T) T {
	return a - b
}

// The comparable Constraint: used when == or != required
func compare[T comparable](a, b T) bool {
	return a == b
}

// Type Sets
type Integer interface { // T can be any if these values
	int8 | int16 | int32 | int64
}

func addGen2[T Integer](a, b T) T {
	return a + b
}

// Using ~
// ex: ~int means int OR any type whose underlying type is int
// func Add[T ~int](a, b T) T {
// 	return a + b
// }
// type MyInt int
// Add(MyInt(1), MyInt(2)) // ✅ works

// Generic structs
type box[T any] struct {
	value T
}

// Generic methods
func (b box[T]) returnValue() T { // refering the above declared struct
	return b.value
}

// Generic Interfaces
type genInteface[T any] interface {
	save(T)
	delete() T
}

// Generic stack
type stack[T any] struct {
	items []T
}

func (s *stack[T]) push(value T) {
	s.items = append(s.items, value)
}

func (s *stack[T]) pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last, true
}

// Generic Map Function
func Map[T any, U any](input []T, f func(T) U) []U {
	result := make([]U, len(input))
	for i, v := range input {
		result[i] = f(v)
	}
	return result
}
