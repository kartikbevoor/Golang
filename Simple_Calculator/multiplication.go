package main

import "fmt"

func mul(nums ...int) {
	if isFirstOp {
		result = nums[0] * nums[1]
		isFirstOp = false
	} else {
		result = result * nums[0]
	}
	fmt.Println("Result:", result)
}
