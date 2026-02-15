package main

import "fmt"

func div(nums ...int) {
	if isFirstOp {
		result = nums[0] / nums[1]
		isFirstOp = false
	} else {
		result = result / nums[0]
	}
	fmt.Println("Result:", result)
}

func isValidDiv(nums ...int) bool {
	if isFirstOp {
		if nums[1] == 0 {
			fmt.Println("Invalid: division by zero is infinity")
			return false
		}
	} else {
		if nums[0] == 0 {
			fmt.Println("Invalid: division by zero is infinity")
			return false
		}
	}
	return true
}
