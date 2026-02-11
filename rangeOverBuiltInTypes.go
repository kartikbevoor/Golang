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

	m := map[string]int{
		"a": 1,
		"b": 2,
	}

	for i, value := range m {
		fmt.Println(i, value)
	}
	// Order is RANDOM, Iteration order is not guaranteed, Changes between runs

}
