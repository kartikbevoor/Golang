package main

import "fmt"

func add(nums ...int) {
	if isFirstOp {
		for _, v := range nums {
			result += v
		}
		isFirstOp = false
	} else {
		for _, v := range nums {
			result = result + v
		}
	}
	fmt.Println("Result:", result)
}
