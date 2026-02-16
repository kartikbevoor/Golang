package main

import "fmt"

// golbal variables
var result int
var isFirstOp bool

func main() {
	isFirstOp = true
	result = 0
	isCalculating := true
	var op string

	fmt.Println("This is a simple calculator.")
	fmt.Println("Enter your choice.")
	fmt.Println(" + for Adition. ")
	fmt.Println(" - for Substraction.")
	fmt.Println(" * for Multiplication.")
	fmt.Println(" / for division.")
	fmt.Println(" = to exit.")

	for isCalculating {
		fmt.Println("Enter your operation: ")
		fmt.Scanln(&op)

		switch op {
		case "+":
			if isFirstOp {
				var a int
				var b int

				fmt.Println("Enter numbers")
				fmt.Scanln(&a)
				fmt.Scanln(&b)

				add(a, b)
			} else {
				var a int
				fmt.Println("Enter number")
				fmt.Scanln(&a)
				add(a)
			}

		case "-":
			if isFirstOp {
				var a int
				var b int

				fmt.Println("Enter numbers")
				fmt.Scanln(&a)
				fmt.Scanln(&b)

				sub(a, b)
			} else {
				var a int
				fmt.Println("Enter number")
				fmt.Scanln(&a)
				sub(a)
			}

		case "*":
			if isFirstOp {
				var a int
				var b int

				fmt.Println("Enter numbers")
				fmt.Scanln(&a)
				fmt.Scanln(&b)

				mul(a, b)
			} else {
				var a int
				fmt.Println("Enter number")
				fmt.Scanln(&a)
				mul(a)
			}

		case "/":
			if isFirstOp {
				var a int
				var b int

				fmt.Println("Enter numbers")
				fmt.Scanln(&a)
				fmt.Scanln(&b)
				if isValidDiv(a, b) {
					div(a, b)
				}
			} else {
				var a int
				fmt.Println("Enter number")
				fmt.Scanln(&a)
				if isValidDiv(a) {
					div(a)
				}
			}

		case "=":
			isCalculating = false
			fmt.Println("Your final result is:", result)
			fmt.Println("Terminated")

		default:
			fmt.Println("Enter valid operation.")
		}
	}

	// div(10, 2)
}
