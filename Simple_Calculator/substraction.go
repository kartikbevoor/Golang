package main

import "fmt"

func sub(nums ...int) {
	if isFirstOp {
		for i, v := range nums {
			if i == 0 {
				result = result + v
			} else {
				result = result - v
			}
		}
		isFirstOp = false
	} else {
		result = result - nums[0]
	}
	fmt.Println("Result:", result)
}
