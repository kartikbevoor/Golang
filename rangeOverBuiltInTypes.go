package main

import "fmt"

// range is a keyword used in for loops to iterate over elements of a collection.
func ranges() {
	nums := [4]int{1, 2, 3, 4}

	// basic syntax
	for i, v := range nums {
		fmt.Println(i, v)
	}

	// use this to exclude index
	// for _, value := range collection

	nums2 := []int{1, 2, 3, 4}

	for _, value := range nums2 {
		fmt.Println(value)
	}

	// note: if u modify value, array does't change

	// this modifies slice elements
	for i := range nums2 {
		nums[i] = nums[i] * 2
	}

	// range over slice
	str := "string"

	for i, r := range str {
		fmt.Println(i, r)
	}

	// range over maps
	m := map[string]int{
		"a": 1,
		"b": 2,
	}

	for i, value := range m {
		fmt.Println(i, value)
	}
	// Order is RANDOM, Iteration order is not guaranteed, Changes between runs

	// cover this topic once you do goroutines and get basic understanding of channesl
	// range over channels
	ch := make(chan int)

	go func() {
		ch <- 1
		ch <- 2
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}

	// Loop continues until channel is CLOSED
	// If not closed → deadlock

	// ignoring index
	for _, value := range nums {
		fmt.Println(value)
	}

	// ignoring value
	for i := range nums {
		fmt.Println(i)
	}

	// Range Copies Values
	for _, v := range nums {
		go func() {
			fmt.Println(v)
		}()
	} // 3, 3, 3

	// to fix the above issue: shadow variable
	for _, value := range nums {
		value := value
		go func() {
			fmt.Println(value)
		}()
	}

	// passing it as parameter
	for _, value := range nums {
		go func(value int) {
			fmt.Println(value)
		}(value)
	}
}
