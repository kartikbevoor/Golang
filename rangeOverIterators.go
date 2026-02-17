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
	for v := range countTo3 {
		if v == 2 {
			break // use of break is to come out of loop
		}
		fmt.Println(v)
	}

	for v := range nums {
		fmt.Println(v)
	}
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
