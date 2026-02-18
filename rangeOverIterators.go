package main

import "fmt"

// An iterator is something that produces values one by one.
// we can use range over function iterators.
// before we could
// for v := range sliceOrChannel {
// }
// now
// for v := range someFunction { // this is ex of range over iterator function
// }

func rangeOverIterators() {
	// Range Over Iterator Functions
	for v := range countTo3 { // here, countTo3 is special function called the iterator function, Where iteratorFunction has a special signature.
		if v == 2 {
			break // use of break is to come out of loop
		}
		fmt.Println(v)
	}

	for v := range nums {
		fmt.Println(v)
	}

	// two value iterator
	for k, v := range pairs {
		fmt.Println(k, v)
	}

	// fibonacci generator
	fiboCount := 0
	fmt.Println("Fibonacci series")
	for v := range fiboGen {
		fmt.Println(v)
		fiboCount++
		if fiboCount == 10 {
			break
		}
	}

	// Using Generics with Iterators
	nums := []int{1, 2, 3}
	for v := range SliceIterator(nums) {
		fmt.Println(v)
	}

	// Tree Traversal Iterator
	// for v := range root.InOrder {
	// 	fmt.Println(v)
	// }

}

// ex
// Range Over Iterator Functions
func countTo3(yield func(int) bool) {
	for i := 0; i < 3; i++ {
		if !yield(i) {
			return
		}
	}
}

func nums(yield func(int) bool) {
	for i := 0; i < 5; i++ {
		if !yield(i) {
			return
		}
	}
}

// what yield function does
// Receives each value, Assigns it to v, Executes the loop body, and Returns true to keep going

// IteratorFunction has a special signature.
// Single value iterator
// func(yield func(T) bool)
// or
// func(func(T) bool)

// Two value iterator
func pairs(yield func(string, int) bool) {
	data := map[string]int{
		"apple":  1,
		"orange": 2,
	}

	for k, v := range data {
		if !yield(k, v) {
			return
		}
	}
}

// ex: fibonacci generator
func fiboGen(yield func(int) bool) {
	a, b := 0, 1
	for {
		if !yield(a) {
			return
		}
		a, b = b, a+b
	}
}

// Using Generics with Iterators
func SliceIterator[T any](slice []T) func(func(T) bool) {
	return func(yield func(T) bool) {
		for _, v := range slice {
			if !yield(v) {
				return
			}
		}
	}
}

// Rules:
// ✔ Must accept a yield function
// ✔ yield must return bool
// ✔ Must stop when yield returns false
// ✔ Can return one or two values
// ✔ No goroutines required

// Tree Traversal Iterator
type Node1 struct {
	Value int
	Left  *Node1
	Right *Node1
}

func (n *Node1) InOrder(yield func(int) bool) bool {
	if n == nil {
		return true
	}

	if !n.Left.InOrder(yield) {
		return false
	}

	if !yield(n.Value) {
		return false
	}

	if !n.Right.InOrder(yield) {
		return false
	}
	return true
}
